package recorders

import "github.com/kira1928/xlive/src/pkg/events"

const (
	RecorderStart   events.EventType = "RecorderStart"
	RecorderStop    events.EventType = "RecorderStop"
	RecorderRestart events.EventType = "RecorderRestart"
)
