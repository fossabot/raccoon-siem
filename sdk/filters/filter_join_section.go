package filters

type filterJoinSection struct {
	or         bool
	not        bool
	conditions []*filterJoinCondition
}

func (s *filterJoinSection) compile(settings *SectionConfig) (*filterJoinSection, error) {
	s.or = settings.Or
	s.not = settings.Not

	for _, expr := range settings.Expressions {
		cond, err := new(filterJoinCondition).compile(expr)

		if err != nil {
			return nil, err
		}

		s.conditions = append(s.conditions, cond)
	}

	return s, nil
}
