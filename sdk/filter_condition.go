package sdk

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	filterConditionExpressionRegexp = regexp.MustCompile(`(\S+)\s(\S+)\s(.+)`)
	FilterIncludeSymbol             = "^"
)

type filterCondition struct {
	incFilterName string
	incFilter     *filter
	leftValue     *smartValue
	operator      byte
	rightValue    *smartValue
}

func (c *filterCondition) compile(expr string) (*filterCondition, error) {
	var err error

	if strings.HasPrefix(expr, FilterIncludeSymbol) {
		c.incFilterName = GetIncludedFilterName(expr)
		return c, nil
	}

	matches := filterConditionExpressionRegexp.FindStringSubmatch(expr)

	if matches == nil {
		return c, fmt.Errorf("invalid filter condition expression '%s", expr)
	}

	// Left value

	leftValueExpr := matches[1]
	c.leftValue = new(smartValue).compile(leftValueExpr)

	// operator

	c.operator, err = ValidateAndTransformFilterOperator(matches[2])

	if err != nil {
		return c, err
	}

	// Right value

	rightValueExpr := matches[3]
	c.rightValue = new(smartValue).compile(rightValueExpr)

	return c, nil
}
