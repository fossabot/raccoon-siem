package normalizers

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	parser "github.com/tephrocactus/raccoon-siem/sdk/normalizers/parsers/regexp"
	"regexp"
)

type regexpNormalizer struct {
	name        string
	expressions []*regexp.Regexp
	mapping     []MappingConfig
}

func (r *regexpNormalizer) ID() string {
	return r.name
}

func (r *regexpNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	parsingResult, ok := parser.Parse(data, r.expressions)
	if !ok || len(parsingResult) == 0 {
		return event
	}
	return normalize(parsingResult, r.mapping, event)
}

func newRegexpNormalizer(cfg Config) (*regexpNormalizer, error) {
	n := &regexpNormalizer{name: cfg.Name}

	for _, e := range cfg.Expressions {
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

		n.expressions = append(n.expressions, compiledExpr)
	}

	return n, nil
}
