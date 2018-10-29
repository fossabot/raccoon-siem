package sdk

type CorrelatorSettings struct {
	DefaultComponentSettings `yaml:",inline"`
	CorrelationRules         []string `yaml:"rules,omitempty"`
	Sources                  []string `yaml:"sources,omitempty"`
	Destinations             []string `yaml:"destinations,omitempty"`
	ActiveListService        string   `yaml:"activeListService,omitempty"`
}

func (s *CorrelatorSettings) ID() string {
	return s.Name
}

type CorrelatorPackage struct {
	DefaultComponentSettings `yaml:",inline"`
	Sources                  []SourceSettings          `yaml:"sources,omitempty"`
	Filters                  []FilterSettings          `yaml:"filters,omitempty"`
	CorrelationRules         []CorrelationRuleSettings `yaml:"rules,omitempty"`
	Destinations             []DestinationSettings     `yaml:"destinations,omitempty"`
	ActiveLists              []ActiveListSettings      `yaml:"activeLists,omitempty"`
}
