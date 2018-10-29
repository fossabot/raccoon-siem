package sdk

import (
	"fmt"
	"regexp"
)

type parserSpecification struct {
	name            string
	kind            string
	subs            []string
	mapping         []*mappingRule
	regexp          *regexp.Regexp
	variants        []*parserVariantSpecification
	variables       map[string]*variable
	rewrites        []*rewriteRule
	kvDelimiter     string
	kvPairDelimiter string
}

func (ps *parserSpecification) compile(settings *ParserSettings) (*parserSpecification, error) {
	// Name

	ps.name = settings.Name

	if ps.name == "" {
		return nil, fmt.Errorf("parser must have a name")
	}

	// Kind

	ps.kind = settings.Kind

	if err := ValidateParserKind(ps.kind); err != nil {
		return nil, err
	}

	// If event parser

	if settings.Kind == parserEvent {
		return ps, nil
	}

	// If regexp parser

	if settings.Kind == parserRegexp {
		ps.regexp = regexp.MustCompile(settings.Regexp)
	}

	// Mapping

	for _, expr := range settings.Mapping {
		mr, err := new(mappingRule).compile(expr, settings.Kind == parserRegexp)

		if err != nil {
			return nil, err
		}

		ps.mapping = append(ps.mapping, mr)
	}

	// Variants

	for _, variant := range settings.Variants {
		vs, err := new(parserVariantSpecification).compile(variant)

		if err != nil {
			return nil, err
		}

		ps.variants = append(ps.variants, vs)
	}

	// Variables
	ps.variables = make(map[string]*variable)

	for _, expr := range settings.Variables {
		vb, err := new(variable).compile(expr)

		if err != nil {
			return nil, err
		}

		if _, ok := ps.variables[vb.name]; ok {
			continue
		}

		ps.variables[vb.name] = vb
	}

	// Rewrites

	for _, expr := range settings.Rewrites {
		rr, err := new(rewriteRule).compile(expr)

		if err != nil {
			return nil, err
		}

		ps.rewrites = append(ps.rewrites, rr)
	}

	// KV Delimiters

	ps.kvDelimiter = settings.KVDelimiter

	if ps.kvDelimiter == "" {
		ps.kvDelimiter = "="
	}

	ps.kvPairDelimiter = settings.KVPairDelimiter

	if ps.kvPairDelimiter == "" {
		ps.kvPairDelimiter = " "
	}

	// Subs

	for _, subName := range settings.Subs {
		if subName == "" {
			return nil, fmt.Errorf("sub parser must have a name")
		}
	}

	ps.subs = settings.Subs

	return ps, nil
}
