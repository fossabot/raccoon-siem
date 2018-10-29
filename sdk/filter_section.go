package sdk

type filterSection struct {
	or         bool
	not        bool
	conditions []*filterCondition
}

func (s *filterSection) compile(settings *FilterSectionSettings) (*filterSection, error) {
	s.or = settings.Or
	s.not = settings.Not

	for _, expr := range settings.Expressions {
		cond, err := new(filterCondition).compile(expr)

		if err != nil {
			return nil, err
		}

		s.conditions = append(s.conditions, cond)
	}

	return s, nil
}
