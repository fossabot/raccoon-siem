package sdk

import (
	"fmt"
	"regexp"
)

var (
	mappingRuleExpressionRegexp = regexp.MustCompile(`(\?)?\s?(\S+)\s?=\s?(\S+)\s?\|?\s?(\w+)?`)
)

type mappingRule struct {
	path       string
	optional   bool
	timeUnit   byte
	eventField string
	index      int
	sub        bool
	variant    bool
}

func (mr *mappingRule) compile(expr string, pathIsIndex bool) (*mappingRule, error) {
	matches := mappingRuleExpressionRegexp.FindStringSubmatch(expr)

	if matches == nil {
		return nil, fmt.Errorf("invalid source field expression '%s'", expr)
	}

	// Optional

	if matches[1] == "?" {
		mr.optional = true
	}

	// Path

	mr.path = matches[3]

	if mr.path == "" {
		return nil, fmt.Errorf("mapping rule path must not be empty")
	}

	if pathIsIndex {
		mr.index = int(toInt(mr.path))
	}

	// Event field

	mr.eventField = matches[2]

	// Field must be sub parsed
	if mr.eventField == "^subs" {
		mr.sub = true
		return mr, nil
	}

	// Field must be parsed in sections
	if mr.eventField == "^variants" {
		mr.variant = true
		return mr, nil
	}

	_, err := ValidateEventFieldAndGetType(mr.eventField)

	if err != nil {
		return nil, err
	}

	// TimeUnit

	timeUnitStr := matches[4]
	mr.timeUnit, err = ValidateAndTransformTimeUnit(timeUnitStr)

	if err != nil {
		return nil, err
	}

	return mr, nil
}
