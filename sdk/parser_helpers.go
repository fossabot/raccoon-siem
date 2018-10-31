package sdk

import (
	"fmt"
)

// Reads configs and registers all parsers
func RegisterParsers(settings []ParserSettings) ([]IParser, error) {
	// Append default Raccoon Event parser
	settings = append(settings, ParserSettings{Name: "event", Kind: "event", Root: true})

	parserSpecs := make([]*parserSpecification, 0, len(settings))
	roots := make(map[string]bool)

	for _, setting := range settings {
		parserSpec, err := setting.Compile()

		if err != nil {
			return nil, err
		}

		roots[setting.Name] = setting.Root
		parserSpecs = append(parserSpecs, parserSpec)
	}

	parsers, err := registerParsers(parserSpecs)

	if err != nil {
		return nil, err
	}

	rootParsers := make([]IParser, 0)

	for _, p := range parsers {
		isRoot := roots[p.ID()]
		if isRoot {
			rootParsers = append(rootParsers, p)
		}
	}

	return rootParsers, nil
}

// Registers and configures all parsers including subs.
func registerParsers(parserSpecs []*parserSpecification) ([]IParser, error) {
	uniqueParsers := make(map[string]IParser)
	result := make([]IParser, 0, len(parserSpecs))

	for _, spec := range parserSpecs {
		if _, ok := uniqueParsers[spec.name]; ok {
			continue
		}

		p := NewParser(spec)
		uniqueParsers[spec.name] = p
		result = append(result, p)
	}

	injectSubs(uniqueParsers)

	return result, nil
}

// Injects subs to any parser (including subs) which requires that
func injectSubs(parsers map[string]IParser) error {
	for _, parser := range parsers {
		for _, subName := range parser.SubNames() {
			sub, ok := parsers[subName]

			if !ok {
				return fmt.Errorf("unknown sub: %s", subName)
			}

			if parser.ID() == sub.ID() {
				return fmt.Errorf("cyclic sub dependency: %s", subName)
			}

			parser.AddSub(sub)
		}
	}
	return nil
}

// Process rewrites
func processRewrites(rewrites []*rewriteRule, variables map[string]*variable, event *Event) {
	for _, rr := range rewrites {
		rr.exec(variables, event)
	}
}
