package logger

import (
	"github.com/tephrocactus/raccoon-siem/sdk"
	"time"
)

const (
	SeverityDebug = "debug"
	LevelDebug    = 1

	SeverityInfo = "info"
	LevelInfo    = 2

	SeverityWarn = "warning"
	LevelWarn    = 3

	SeverityError = "error"
	LevelError    = 4
)

type Instance struct {
	name         string
	level        int
	destinations []sdk.IDestination
}

func (r *Instance) Debug(msg string, customEvent ...*sdk.Event) {
	if r.level <= LevelDebug {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = SeverityDebug
		r.sendEvent(e)
	}
}

func (r *Instance) Info(msg string, customEvent ...*sdk.Event) {
	if r.level <= LevelInfo {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = SeverityInfo
		r.sendEvent(e)
	}
}

func (r *Instance) Warn(msg string, customEvent ...*sdk.Event) {
	if r.level <= LevelWarn {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = SeverityWarn
		r.sendEvent(e)
	}
}

func (r *Instance) Error(msg string, customEvent ...*sdk.Event) {
	if r.level <= LevelError {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = SeverityError
		r.sendEvent(e)
	}
}

func (r *Instance) sendEvent(e *sdk.Event) {
	for _, dst := range r.destinations {
		dst.Send(e)
	}
}

func (r *Instance) defaultEvent(msg string, customEvent ...*sdk.Event) *sdk.Event {
	var e *sdk.Event

	if len(customEvent) > 0 {
		e = customEvent[0]
	} else {
		e = new(sdk.Event)
	}

	e.OriginTimestamp = time.Now()
	e.OriginServiceName = r.name
	e.Message = msg

	return e
}

func NewInstance(name string, destinations []sdk.IDestination, level ...int) *Instance {
	effectiveLogLevel := LevelWarn
	if len(level) != 0 {
		effectiveLogLevel = level[0]
	}

	return &Instance{
		name:         name,
		level:        effectiveLogLevel,
		destinations: destinations,
	}
}
