package regexp

import (
	"fmt"
	"regexp"
)

type Parser struct {
	name        string
	expressions []*regexp.Regexp
}

func (r *Parser) ID() string {
	return r.name
}

func (r *Parser) Parse(data []byte) (map[string]string, bool) {
	for _, e := range r.expressions {
		if match := e.FindSubmatch(data); match != nil {
			output := make(map[string]string)
			for i, field := range e.SubexpNames() {
				if i > 0 {
					output[field] = string(match[i])
				}
			}
			return output, true
		}
	}
	return nil, false
}

func NewParser(name string, expressions []string) (*Parser, error) {
	p := &Parser{name: name}
	for _, e := range expressions {
		compiledExpr, err := regexp.Compile(e)
		if err != nil {
			return nil, err
		}

		groupNum := compiledExpr.NumSubexp()
		if groupNum == 0 {
			return nil, fmt.Errorf("at least one capturing group must be specified: %s", e)
		}

		names := compiledExpr.SubexpNames()
		for i := 1; i < len(names); i++ {
			if names[i] == "" {
				return nil, fmt.Errorf("each capturing group must have a name: %s", e)
			}
		}

		p.expressions = append(p.expressions, compiledExpr)
	}
	return p, nil
}
