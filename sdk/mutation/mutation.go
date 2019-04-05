package mutation

import (
	"regexp"
	"strings"
)

func Mutate(configs []Config, value string) string {
	for _, cfg := range configs {
		value = mutate(cfg, value)
	}
	return value
}

func mutate(cfg Config, value string) string {
	switch cfg.Kind {
	case KindRegexp:
		return mutateRegexp(cfg.expression, value)
	case KindLower:
		return mutateLower(value)
	case KindSubstring:
		return mutateSubstring(value, cfg.Start, cfg.End)
	}
	return value
}

func mutateRegexp(expr *regexp.Regexp, value string) string {
	if match := expr.FindStringSubmatch(value); len(match) > 1 {
		return match[1]
	}
	return value
}

func mutateLower(value string) string {
	return strings.ToLower(value)
}

func mutateSubstring(value string, start, end int) string {
	if start < len(value) && end <= len(value) {
		return value[start:end]
	}
	return value
}
