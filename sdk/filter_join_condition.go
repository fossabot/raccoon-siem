package sdk

import (
	"fmt"
	"regexp"
)

var (
	filterJoinConditionExpressionRegexp = regexp.MustCompile(`([^\.]+)\.([^\.]+)\s(\S+)\s([^\.]+)\.([^\.]+)`)
)

type filterJoinCondition struct {
	LeftEventID     string
	LeftEventField  string
	Operator        byte
	RightEventID    string
	RightEventField string
}

func (c *filterJoinCondition) compile(expr string) (*filterJoinCondition, error) {
	var err error

	matches := filterJoinConditionExpressionRegexp.FindStringSubmatch(expr)

	if matches == nil {
		return c, fmt.Errorf("invalid join filter condition expression '%s", expr)
	}

	// Left operand

	c.LeftEventID = matches[1]
	c.LeftEventField = matches[2]

	if _, err := ValidateEventFieldAndGetType(c.LeftEventField); err != nil {
		return c, err
	}

	// operator

	c.Operator, err = ValidateAndTransformFilterOperator(matches[3])

	if err != nil {
		return c, err
	}

	// Right operand

	c.RightEventID = matches[4]
	c.RightEventField = matches[5]

	if _, err := ValidateEventFieldAndGetType(c.RightEventField); err != nil {
		return c, err
	}

	return c, nil
}
