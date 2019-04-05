package mutation

import (
	"fmt"
	"regexp"
)

const (
	KindRegexp = "regexp"
)

type Config struct {
	Kind       string         `json:"kind"`
	Expression string         `json:"expression"`
	expression *regexp.Regexp `json:"-"`
}

func (r *Config) Validate() (err error) {
	switch r.Kind {
	case KindRegexp:
	default:
		return fmt.Errorf("mutation: invalid kind %s", r.Kind)
	}

	r.expression, err = regexp.Compile(r.Expression)
	if err != nil {
		return err
	}

	return nil
}
