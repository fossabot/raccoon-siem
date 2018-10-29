package sdk

var knownTriggers = map[string]bool{
	triggerEveryEvent:           true,
	triggerFirstEvent:           true,
	triggerSubsequentEvents:     true,
	triggerEveryThreshold:       true,
	triggerFirstThreshold:       true,
	triggerSubsequentThresholds: true,
	triggerAllThresholdsReached: true,
	triggerTimeout:              true,
}

type TriggerSettings struct {
	Kind    string           `yaml:"kind,omitempty"`
	Actions []ActionSettings `yaml:"actions,omitempty"`
}

func (s *TriggerSettings) compile() (string, []*actionSpecifications, error) {
	acts := make([]*actionSpecifications, 0)

	// Kind

	if err := ValidateTrigger(s.Kind); err != nil {
		return "", nil, err
	}

	// Actions

	for _, actionSetting := range s.Actions {
		action, err := actionSetting.compile()

		if err != nil {
			return "", nil, err
		}

		acts = append(acts, action)
	}

	return s.Kind, acts, nil
}
