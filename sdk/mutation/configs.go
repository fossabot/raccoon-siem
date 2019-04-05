package mutation

import (
	"fmt"
	"regexp"
)

const (
	KindRegexp    = "regexp"
	KindLower     = "lower"
	KindSubstring = "substring"
)

type Config struct {
	Kind       string         `json:"kind"`
	Expression string         `json:"expression"`
	Start      int            `json:"start"`
	End        int            `json:"end"`
	expression *regexp.Regexp `json:"-"`
}

func (r *Config) Validate() (err error) {
	switch r.Kind {
	case KindRegexp, KindLower, KindSubstring:
	default:
		return fmt.Errorf("mutation: invalid kind %s", r.Kind)
	}

	r.expression, err = regexp.Compile(r.Expression)
	if err != nil {
		return err
	}

	if r.Start < 0 || r.End < 0 {
		return fmt.Errorf("mutation: start and end must be positive numbers")
	}

	return nil
}
