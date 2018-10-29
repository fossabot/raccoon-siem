package sdk

import (
	"fmt"
)

const (
	parserJSON   = "json"
	parserEvent  = "event"
	parserRegexp = "regexp"
	parserSyslog = "syslog"
	parserKV     = "kv"
)

var knownParserKinds = map[string]bool{
	parserJSON:   true,
	parserEvent:  true,
	parserRegexp: true,
	parserSyslog: true,
	parserKV:     true,
}

type IParser interface {
	ID() string
	Parse(data []byte, target *Event) (*Event, error)
	AddSub(sub IParser)
	SubNames() []string
}

func NewParser(spec *parserSpecification) IParser {
	switch spec.kind {
	case parserJSON:
		return newJSONParser(spec)
	case parserEvent:
		return newEventParser(spec)
	case parserSyslog:
		return newSyslogParser(spec)
	case parserRegexp:
		return newRegexpParser(spec)
	case parserKV:
		return newKeyValParser(spec)
	default:
		panic(fmt.Errorf("unknown parser type: %s", spec.kind))
	}
}
