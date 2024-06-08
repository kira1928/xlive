package tools

import (
	"fmt"

	"github.com/kira1928/remotetools"
	_tools "github.com/kira1928/remotetools/pkg/tools"
	"github.com/kira1928/xlive/src/log"
)

var (
	instance  *_tools.API
	toolNames = []string{
		"dotnet",
		"ffmpeg",
		"brec",
	}
)

func init() {
	_tools.SetToolFolder("external_tools")
	instance = remotetools.Get()
}

func Init() (err error) {
	if err = instance.LoadConfig("remote_tools.json"); err != nil {
		return
	}
	for _, toolName := range toolNames {
		if err = initImpl(toolName); err != nil {
			return
		}
	}
	return
}

func initImpl(toolName string) (err error) {
	tool, err := instance.GetTool(toolName)
	if err != nil {
		return
	}
	if tool.DoesToolExist() {
		log.Get().Infof("[tool]%s (version %s) is ready", toolName, tool.GetVersion())
	} else {
		log.Get().Infof("[tool]installing %s (version %s) ...", toolName, tool.GetVersion())
		if err = tool.Install(); err != nil {
			return fmt.Errorf("[tool]install faliled: %s", err.Error())
		}
		log.Get().Infof("[tool]%s (version %s) is installed", toolName, tool.GetVersion())
	}
	return
}
