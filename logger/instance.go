package logger

import (
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
)

type Instance struct {
	name         string
	level        int
	destinations []destinations.IDestination
}

func (r *Instance) Debug(msg string, customEvent ...*normalization.Event) {
	if r.level <= LevelDebug {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = normalization.SeverityInfo
		r.sendEvent(e)
	}
}

func (r *Instance) Info(msg string, customEvent ...*normalization.Event) {
	if r.level <= LevelInfo {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = normalization.SeverityInfo
		r.sendEvent(e)
	}
}

func (r *Instance) Warn(msg string, customEvent ...*normalization.Event) {
	if r.level <= LevelWarn {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = normalization.SeverityWarn
		r.sendEvent(e)
	}
}

func (r *Instance) Error(msg string, customEvent ...*normalization.Event) {
	if r.level <= LevelError {
		e := r.defaultEvent(msg, customEvent...)
		e.Severity = normalization.SeverityError
		r.sendEvent(e)
	}
}

func (r *Instance) sendEvent(e *normalization.Event) {
	for _, dst := range r.destinations {
		dst.Send(e)
	}
}

func (r *Instance) defaultEvent(msg string, customEvent ...*normalization.Event) *normalization.Event {
	var e *normalization.Event

	if len(customEvent) > 0 {
		e = customEvent[0]
	} else {
		e = new(normalization.Event)
	}

	e.OriginTimestamp = helpers.NowUnixMillis()
	e.OriginServiceName = r.name
	e.Message = msg

	return e
}

func NewInstance(name string, destinations []destinations.IDestination, level ...int) *Instance {
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
