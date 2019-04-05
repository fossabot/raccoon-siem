package mutation

func Mutate(configs []Config, value string) string {
	for _, cfg := range configs {
		value = mutate(cfg, value)
	}
	return value
}

func mutate(cfg Config, value string) string {
	switch cfg.Kind {
	case KindRegexp:
		return mutateRegexp(cfg, value)
	}
	return value
}

func mutateRegexp(cfg Config, value string) string {
	if match := cfg.expression.FindStringSubmatch(value); len(match) > 1 {
		return match[1]
	}
	return value
}
