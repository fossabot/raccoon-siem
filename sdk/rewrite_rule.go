package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"regexp"
)

var (
	rewriteRuleExpressionRegexp = regexp.MustCompile(`(\S+) = (.+)`)
)

type rewriteRule struct {
	targetEventField string
	sourceValue      *smartValue
}

func (r *rewriteRule) compile(expr string) (*rewriteRule, error) {
	matches := rewriteRuleExpressionRegexp.FindStringSubmatch(expr)

	if matches == nil {
		return nil, fmt.Errorf("invalid rewrite rule expression '%s'", expr)
	}

	r.targetEventField = matches[1]

	if _, err := ValidateEventFieldAndGetType(r.targetEventField); err != nil {
		return nil, err
	}

	r.sourceValue = new(smartValue).compile(matches[2])

	return r, nil
}

func (r *rewriteRule) exec(vars map[string]*variable, event *normalization.Event) {
	v, err := r.sourceValue.resolve(vars, event)

	if err != nil {
		if Debug {
			fmt.Println(err)
		}
		return
	}

	event.SetField(r.targetEventField, v, normalization.TimeUnitNone)
}
