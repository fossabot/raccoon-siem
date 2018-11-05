package sdk

import (
	"fmt"
	"github.com/jeromer/syslogparser/rfc3164"
)

func newSyslogParser(spec *parserSpecification) IParser {
	return &syslogParser{spec: spec}
}

type syslogParser struct {
	spec *parserSpecification
	subs []IParser
}

func (p *syslogParser) AddSub(sub IParser) {
	p.subs = append(p.subs, sub)
}

func (p *syslogParser) SubNames() []string {
	return p.spec.subs
}

func (p *syslogParser) ID() string {
	return p.spec.name
}

func (p *syslogParser) Parse(input []byte, target *Event) (*Event, error) {
	event := target

	if event == nil {
		event = new(Event)
	}

	sp := rfc3164.NewParser(input)
	if err := sp.Parse(); err != nil {
		return nil, err
	}

	// timestamp, hostname, tag, content, priority, facility, severity
	syslogFieldsMap := sp.Dump()

	for _, field := range p.spec.mapping {
		parsedValue, ok := syslogFieldsMap[field.path]

		if !ok {
			if !field.optional {
				return nil, fmt.Errorf("syslog field '%s' not found", field.path)
			}
			continue
		}

		if field.sub {
			if err := p.subParse([]byte(parsedValue.(string)), field, event); err != nil {
				if !field.optional {
					return nil, err
				}
			}
			continue
		}

		event.SetField(field.eventField, parsedValue, field.timeUnit)
	}

	if len(p.spec.rewrites) != 0 {
		processRewrites(p.spec.rewrites, p.spec.variables, event)
	}

	return event, nil
}

func (p *syslogParser) subParse(input []byte, field *mappingRule, target *Event) error {
	subParsed := false

	for _, sub := range p.subs {
		if _, err := sub.Parse(input, target); err != nil {
			continue
		}

		subParsed = true
		break
	}

	if !subParsed {
		return ErrAllParsersFailed
	}

	return nil
}
