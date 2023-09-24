//go:generate mockgen -package recorders -destination mock_test.go github.com/hr3lxphr6j/bililive-go/src/recorders Recorder,Manager
package recorders

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"

	"github.com/bluele/gcache"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/hr3lxphr6j/bililive-go/src/configs"
	"github.com/hr3lxphr6j/bililive-go/src/instance"
	"github.com/hr3lxphr6j/bililive-go/src/interfaces"
	"github.com/hr3lxphr6j/bililive-go/src/live"
	"github.com/hr3lxphr6j/bililive-go/src/live/bilibili"
	"github.com/hr3lxphr6j/bililive-go/src/live/douyu"
	"github.com/hr3lxphr6j/bililive-go/src/live/huya"
	"github.com/hr3lxphr6j/bililive-go/src/pkg/events"
	"github.com/hr3lxphr6j/bililive-go/src/pkg/parser"
	"github.com/hr3lxphr6j/bililive-go/src/pkg/parser/ffmpeg"
	"github.com/hr3lxphr6j/bililive-go/src/pkg/parser/native/flv"
	"github.com/hr3lxphr6j/bililive-go/src/pkg/utils"
)

const (
	begin uint32 = iota
	pending
	running
	stopped
)

// for test
var (
	newParser = func(u *url.URL, useNativeFlvParser bool, cfg map[string]string) (parser.Parser, error) {
		parserName := ffmpeg.Name
		if strings.Contains(u.Path, ".flv") && useNativeFlvParser {
			parserName = flv.Name
		}
		return parser.New(parserName, cfg)
	}

	mkdir = func(path string) error {
		return os.MkdirAll(path, os.ModePerm)
	}

	removeEmptyFile = func(file string) {
		if stat, err := os.Stat(file); err == nil && stat.Size() == 0 {
			os.Remove(file)
		}
	}
)

func getDefaultFileNameTmpl(config *configs.Config) *template.Template {
	return template.Must(template.New("filename").Funcs(utils.GetFuncMap(config)).
		Parse(`{{ .Live.GetPlatformCNName }}/{{ .HostName | filenameFilter }}/[{{ now | date "2006-01-02 15-04-05"}}][{{ .HostName | filenameFilter }}][{{ .RoomName | filenameFilter }}].flv`))
}

type Recorder interface {
	Start(ctx context.Context) error
	StartTime() time.Time
	GetStatus() (map[string]string, error)
	Close()
}

type recorder struct {
	Live       live.Live
	OutPutPath string

	config     *configs.Config
	ed         events.Dispatcher
	logger     *interfaces.Logger
	cache      gcache.Cache
	startTime  time.Time
	parser     parser.Parser
	parserLock *sync.RWMutex

	stop      chan struct{}
	stopDanmu chan struct{}
	state     uint32
}

func NewRecorder(ctx context.Context, live live.Live) (Recorder, error) {
	inst := instance.GetInstance(ctx)
	return &recorder{
		Live:       live,
		OutPutPath: instance.GetInstance(ctx).Config.OutPutPath,
		config:     inst.Config,
		cache:      inst.Cache,
		startTime:  time.Now(),
		ed:         inst.EventDispatcher.(events.Dispatcher),
		logger:     inst.Logger,
		state:      begin,
		stop:       make(chan struct{}),
		stopDanmu:  make(chan struct{}),
		parserLock: new(sync.RWMutex),
	}, nil
}

func (r *recorder) tryRecord(ctx context.Context) {
	urls, err := r.Live.GetStreamUrls()
	if err != nil || len(urls) == 0 {
		r.getLogger().WithError(err).Warn("failed to get stream url, will retry after 5s...")
		time.Sleep(5 * time.Second)
		return
	}

	obj, _ := r.cache.Get(r.Live)
	info := obj.(*live.Info)

	tmpl := getDefaultFileNameTmpl(r.config)
	if r.config.OutputTmpl != "" {
		_tmpl, err := template.New("user_filename").Funcs(utils.GetFuncMap(r.config)).Parse(r.config.OutputTmpl)
		if err == nil {
			tmpl = _tmpl
		}
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, info); err != nil {
		panic(fmt.Sprintf("failed to render filename, err: %v", err))
	}
	fileName := filepath.Join(r.OutPutPath, buf.String())
	outputPath, _ := filepath.Split(fileName)
	url := urls[0]

	if strings.Contains(url.Path, "m3u8") {
		fileName = fileName[:len(fileName)-4] + ".ts"
	}

	if info.AudioOnly {
		fileName = fileName[:strings.LastIndex(fileName, ".")] + ".aac"
	}

	if err = mkdir(outputPath); err != nil {
		r.getLogger().WithError(err).Errorf("failed to create output path[%s]", outputPath)
		return
	}
	parserCfg := map[string]string{
		"timeout_in_us": strconv.Itoa(r.config.TimeoutInUs),
	}
	if r.config.Debug {
		parserCfg["debug"] = "true"
	}
	p, err := newParser(url, r.config.Feature.UseNativeFlvParser, parserCfg)
	if err != nil {
		r.getLogger().WithError(err).Error("failed to init parse")
		return
	}
	r.setAndCloseParser(p)
	r.startTime = time.Now()
	r.getLogger().Debugln("Start ParseLiveStream(" + url.String() + ", " + fileName + ")")
	go r.runDanmu(ctx, fileName)
	r.getLogger().Println(r.parser.ParseLiveStream(ctx, url, r.Live, fileName))
	close(r.stopDanmu)
	r.getLogger().Debugln("End ParseLiveStream(" + url.String() + ", " + fileName + ")")
	removeEmptyFile(fileName)
	ffmpegPath, err := utils.GetFFmpegPath(ctx)
	if err != nil {
		r.getLogger().WithError(err).Error("failed to find ffmpeg")
		return
	}
	cmdStr := strings.Trim(r.config.OnRecordFinished.CustomCommandline, "")
	if len(cmdStr) > 0 {
		tmpl, err := template.New("custom_commandline").Funcs(utils.GetFuncMap(r.config)).Parse(cmdStr)
		if err != nil {
			r.getLogger().WithError(err).Error("custom commandline parse failure")
			return
		}

		obj, _ := r.cache.Get(r.Live)
		info := obj.(*live.Info)

		buf := new(bytes.Buffer)
		if err := tmpl.Execute(buf, struct {
			*live.Info
			FileName string
			Ffmpeg   string
		}{
			Info:     info,
			FileName: fileName,
			Ffmpeg:   ffmpegPath,
		}); err != nil {
			r.getLogger().WithError(err).Errorln("failed to render custom commandline")
			return
		}
		bash := ""
		args := []string{}
		switch runtime.GOOS {
		case "linux":
			bash = "sh"
			args = []string{"-c"}
		case "windows":
			bash = "cmd"
			args = []string{"/C"}
		default:
			r.getLogger().Warnln("Unsupport system ", runtime.GOOS)
		}
		args = append(args, buf.String())
		r.getLogger().Debugf("start executing custom_commandline: %s", args[1])
		cmd := exec.Command(bash, args...)
		if r.config.Debug {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		if err = cmd.Run(); err != nil {
			r.getLogger().WithError(err).Debugf("custom commandline execute failure (%s %s)\n", bash, strings.Join(args, " "))
		} else if r.config.OnRecordFinished.DeleteFlvAfterConvert {
			os.Remove(fileName)
		}
		r.getLogger().Debugf("end executing custom_commandline: %s", args[1])
	} else if r.config.OnRecordFinished.ConvertToMp4 {
		convertCmd := exec.Command(
			ffmpegPath,
			"-hide_banner",
			"-i",
			fileName,
			"-c",
			"copy",
			fileName+".mp4",
		)
		if err = convertCmd.Run(); err != nil {
			convertCmd.Process.Kill()
			r.getLogger().Debugln(err)
		} else if r.config.OnRecordFinished.DeleteFlvAfterConvert {
			os.Remove(fileName)
		}
	}
}

func (r *recorder) run(ctx context.Context) {
	for {
		select {
		case <-r.stop:
			return
		default:
			r.tryRecord(ctx)
		}
	}
}

type BadanmuResponseUserInfo struct {
	Gender   int
	UserId   int
	UserName string
}
type BadanmuResponse struct {
	CommonType int
	Data       string
	RoomId     string
	Type       string
	Code       int
	Ts         int
	PlayerName string
	UserInfo   BadanmuResponseUserInfo
}

func (r *recorder) runDanmu(ctx context.Context, fileName string) {
	obj, _ := r.cache.Get(r.Live)
	info := obj.(*live.Info)
	platformCnName := info.Live.GetPlatformCNName()
	platform := ""
	roomId := ""
	switch platformCnName {
	case huya.CnName:
		platform = "huya"
		roomId = info.Live.(*huya.Live).RoomId
	case douyu.CnName:
		platform = "douyu"
		roomId = info.Live.(*douyu.Live).RoomID
	case bilibili.CnName:
		platform = "bilibili"
		roomId = info.Live.(*bilibili.Live).RealID
	default:
		// todo error log
		return
	}

	// todo customize port in config
	url := "ws://badanmu:8181/"
	c, _, err := websocket.Dial(ctx, url, &websocket.DialOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer c.Close(websocket.StatusInternalError, "runDanmu() finished")

	v := BadanmuResponse{}
	var db *sql.DB
	r.stopDanmu = make(chan struct{})
	for {
		select {
		case <-r.stopDanmu:
			return
		default:
			err := wsjson.Read(ctx, c, &v)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			switch v.CommonType {
			case 1001:
				if v.Data == "success" {
					db, err = utils.CreateDanmuFile(strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".sqlite")
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				} else {
					err = wsjson.Write(ctx, c, map[string]string{
						"type":     "login",
						"platform": platform,
						"roomId":   roomId,
					})
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}
			case 0:
				switch v.Type {
				case "comment":
					err := utils.AddDanmuRecord(db, v.PlayerName, v.Data, v.UserInfo.UserId, v.Ts)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}
			}
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(b))
		}
	}

}

func (r *recorder) getParser() parser.Parser {
	r.parserLock.RLock()
	defer r.parserLock.RUnlock()
	return r.parser
}

func (r *recorder) setAndCloseParser(p parser.Parser) {
	r.parserLock.Lock()
	defer r.parserLock.Unlock()
	if r.parser != nil {
		r.parser.Stop()
	}
	r.parser = p
}

func (r *recorder) Start(ctx context.Context) error {
	if !atomic.CompareAndSwapUint32(&r.state, begin, pending) {
		return nil
	}
	go r.run(ctx)
	r.getLogger().Info("Record Start")
	r.ed.DispatchEvent(events.NewEvent(RecorderStart, r.Live))
	atomic.CompareAndSwapUint32(&r.state, pending, running)
	return nil
}

func (r *recorder) StartTime() time.Time {
	return r.startTime
}

func (r *recorder) Close() {
	if !atomic.CompareAndSwapUint32(&r.state, running, stopped) {
		return
	}
	close(r.stop)
	if p := r.getParser(); p != nil {
		p.Stop()
	}
	r.getLogger().Info("Record End")
	r.ed.DispatchEvent(events.NewEvent(RecorderStop, r.Live))
}

func (r *recorder) getLogger() *logrus.Entry {
	return r.logger.WithFields(r.getFields())
}

func (r *recorder) getFields() map[string]interface{} {
	obj, err := r.cache.Get(r.Live)
	if err != nil {
		return nil
	}
	info := obj.(*live.Info)
	return map[string]interface{}{
		"host": info.HostName,
		"room": info.RoomName,
	}
}

func (r *recorder) GetStatus() (map[string]string, error) {
	statusP, ok := r.getParser().(parser.StatusParser)
	if !ok {
		return nil, ErrParserNotSupportStatus
	}
	return statusP.Status()
}
