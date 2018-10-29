package sdk

import (
	"github.com/pkg/errors"
)

var errRegexpNoMatch = errors.New("nothing matched")

func newRegexpParser(spec *parserSpecification) IParser {
	return &regexpParser{spec: spec}
}

type regexpParser struct {
	spec *parserSpecification
	subs []IParser
}

func (p *regexpParser) AddSub(sub IParser) {
	p.subs = append(p.subs, sub)
}

func (p *regexpParser) SubNames() []string {
	return p.spec.subs
}

func (p *regexpParser) ID() string {
	return p.spec.name
}

func (p *regexpParser) Parse(input []byte, target *Event) (*Event, error) {
	event := target

	if event == nil {
		event = new(Event)
	}

	// Match regexp

	matches := p.spec.regexp.FindStringSubmatch(string(input))

	if matches == nil {
		return nil, errRegexpNoMatch
	}

	// Run through mapping rules

	for _, mappingRule := range p.spec.mapping {
		parsedValue := matches[mappingRule.index]

		// Continue parsing with variants

		if mappingRule.variant {
			if err := p.variantParse(parsedValue, event); err != nil {
				if !mappingRule.optional {
					return nil, err
				}
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

	return event, nil
}

func (p *regexpParser) variantParse(input string, event *Event) error {
	parsed := false

	for _, v := range p.spec.variants {
		matches := v.regexp.FindStringSubmatch(input)

		if matches == nil {
			continue
		}

		for _, field := range v.mapping {
			parsedValue := matches[field.index]
			event.SetField(field.eventField, parsedValue, field.timeUnit)
		}

		parsed = true
		break
	}

	if !parsed {
		return ErrAllParsersFailed
	}

	return nil
}

func (p *regexpParser) subParse(input []byte, field *mappingRule, target *Event) error {
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
