package sdk

import (
	jp "github.com/tidwall/gjson"
)

func newJSONParser(spec *parserSpecification) IParser {
	return &jsonParser{spec: spec}
}

type jsonParser struct {
	spec *parserSpecification
	subs []IParser
}

func (p *jsonParser) AddSub(sub IParser) {
	p.subs = append(p.subs, sub)
}

func (p *jsonParser) SubNames() []string {
	return p.spec.subs
}

func (p *jsonParser) ID() string {
	return p.spec.name
}

func (p *jsonParser) Parse(input []byte, target *Event) (*Event, error) {
	event := target

	if event == nil {
		event = new(Event)
	}

	// Run through mapping rules

	for _, mappingRule := range p.spec.mapping {
		parsedValue, err := p.parseField(input, mappingRule.path)

		if err != nil {
			if !mappingRule.optional {
				return nil, err
			}
			continue
		}

		if mappingRule.sub {
			if err := p.subParse(input, mappingRule, event); err != nil {
				if !mappingRule.optional {
					return nil, err
				}
			}
			continue
		}

		event.SetField(mappingRule.eventField, parsedValue, mappingRule.timeUnit)
	}

	if len(p.spec.rewrites) != 0 {
		processRewrites(p.spec.rewrites, p.spec.variables, event)
	}

	return event, nil
}

func (p *jsonParser) subParse(input []byte, mr *mappingRule, target *Event) error {
	subParsed := false
	subValue := jp.GetBytes(input, mr.path)

	if !subValue.Exists() {
		return ErrJSONFieldDoesNotExist
	}

	subBytes := []byte(subValue.Raw)

	for _, sub := range p.subs {
		if _, err := sub.Parse(subBytes, target); err != nil {
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

func (p *jsonParser) parseField(input []byte, path string) (parsedValue interface{}, err error) {
	value := jp.GetBytes(input, path)

	if !value.Exists() {
		return nil, ErrJSONFieldDoesNotExist
	}

	switch value.Type {
	case jp.String:
		parsedValue = value.String()
	case jp.Number:
		parsedValue = value.Float()
	case jp.False:
		fallthrough
	case jp.True:
		parsedValue = value.Bool()
	case jp.Null:
		parsedValue = nil
	default:
		parsedValue = value.Raw
	}

	return parsedValue, err
}
