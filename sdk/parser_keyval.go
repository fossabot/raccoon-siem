package sdk

import (
	"fmt"
	"strings"
)

func newKeyValParser(spec *parserSpecification) IParser {
	return &keyValParser{
		spec: spec,
	}
}

type keyValParser struct {
	spec *parserSpecification
	subs []IParser
}

func (p *keyValParser) AddSub(sub IParser) {
	p.subs = append(p.subs, sub)
}

func (p *keyValParser) SubNames() []string {
	return p.spec.subs
}

func (p *keyValParser) ID() string {
	return p.spec.name
}

func (p *keyValParser) Parse(input []byte, target *Event) (*Event, error) {
	event := target

	if event == nil {
		event = new(Event)
	}

	// TODO: work with byte slice

	pairs := strings.Split(string(input), p.spec.kvPairDelimiter)

	m := make(map[string]string)

	for i := range pairs {
		if pairs[i] == "" {
			continue
		}

		kv := strings.Split(pairs[i], p.spec.kvDelimiter)

		if len(kv) != 2 {
			return nil, errMalformedInput
		}

		m[kv[0]] = kv[1]
	}

	for _, mappingRule := range p.spec.mapping {
		parsedValue, ok := m[mappingRule.path]

		if !ok {
			if !mappingRule.optional {
				return nil, errKeyNotFound
			}
			continue
		}

		// Continue parsing with sub parsers

		if mappingRule.sub {
			if err := p.subParse([]byte(parsedValue), mappingRule, event); err != nil {
				if !mappingRule.optional {
					return nil, err
				}
			}
			continue
		}

		// Set event field value

		event.SetField(mappingRule.eventField, parsedValue, mappingRule.timeUnit)
	}

	if len(p.spec.rewrites) != 0 {
		processRewrites(p.spec.rewrites, p.spec.variables, event)
	}

	if event.OriginServiceName != "netflow" {
		fmt.Println(event)
	}

	return event, nil
}

func (p *keyValParser) subParse(input []byte, field *mappingRule, target *Event) error {
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
