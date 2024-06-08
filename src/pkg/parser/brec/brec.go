package brec

import (
	"context"
	"io"
	"net/url"
	"os"
	"os/exec"
	"sync"

	"github.com/kira1928/remotetools"
	"github.com/kira1928/xlive/src/live"
	"github.com/kira1928/xlive/src/pkg/parser"
)

const (
	Name = "brec"

	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
)

func init() {
	parser.Register(Name, new(builder))
}

type builder struct{}

func (b *builder) Build(cfg map[string]string) (parser.Parser, error) {
	debug := false
	if debugFlag, ok := cfg["debug"]; ok && debugFlag != "" {
		debug = true
	}
	return &Parser{
		debug:       debug,
		closeOnce:   new(sync.Once),
		timeoutInUs: cfg["timeout_in_us"],
	}, nil
}

type Parser struct {
	cmd         *exec.Cmd
	cmdStdIn    io.WriteCloser
	cmdStdout   io.ReadCloser
	closeOnce   *sync.Once
	debug       bool
	timeoutInUs string

	cmdLock sync.Mutex
}

func (p *Parser) ParseLiveStream(ctx context.Context, url *url.URL, live live.Live, file string) (err error) {
	dotnet, err := remotetools.Get().GetTool("dotnet")
	if err != nil {
		return err
	}
	brec, err := remotetools.Get().GetTool("brec")
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	headers := live.GetHeadersForDownloader()
	_, exists := headers["User-Agent"]
	if !exists {
		headers["User-Agent"] = userAgent
	}
	_, exists = headers["Referer"]
	if !exists {
		headers["Referer"] = live.GetRawUrl()
	}
	args := []string{
		brec.GetToolPath(),
		"d",
	}
	for k, v := range headers {
		args = append(args, "-h", k+": "+v)
	}
	args = append(args, url.String())
	args = append(args, file)

	// p.cmd operations need p.cmdLock
	func() {
		p.cmdLock.Lock()
		defer p.cmdLock.Unlock()
		p.cmd = exec.Command(dotnet.GetToolPath(), args...)
		if p.cmdStdIn, err = p.cmd.StdinPipe(); err != nil {
			return
		}
		if p.cmdStdout, err = p.cmd.StdoutPipe(); err != nil {
			return
		}
		if p.debug {
			p.cmd.Stderr = os.Stderr
		}
		if err = p.cmd.Start(); err != nil {
			p.cmd.Process.Kill()
			return
		}
	}()
	if err != nil {
		return err
	}

	err = p.cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) Stop() (err error) {
	p.closeOnce.Do(func() {
		p.cmdLock.Lock()
		defer p.cmdLock.Unlock()
		if p.cmd != nil && p.cmd.ProcessState == nil {
			if p.cmd.Process != nil {
				err = p.cmd.Process.Kill()
			}
		}
	})
	return
}
