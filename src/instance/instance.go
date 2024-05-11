package instance

import (
	"sync"

	"github.com/bluele/gcache"

	"github.com/kira1928/xlive/src/configs"
	"github.com/kira1928/xlive/src/interfaces"
	"github.com/kira1928/xlive/src/live"
)

type Instance struct {
	WaitGroup       sync.WaitGroup
	Config          *configs.Config
	Logger          *interfaces.Logger
	Lives           map[live.ID]live.Live
	Cache           gcache.Cache
	Server          interfaces.Module
	EventDispatcher interfaces.Module
	ListenerManager interfaces.Module
	RecorderManager interfaces.Module
}
