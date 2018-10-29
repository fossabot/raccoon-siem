package sdk

type FilterSettings struct {
	Name      string                  `yaml:"name,omitempty"`
	Not       bool                    `yaml:"not,omitempty"`
	Join      bool                    `yaml:"join,omitempty"`
	Sections  []FilterSectionSettings `yaml:"sections,omitempty"`
	Variables []string                `yaml:"vars,omitempty"`
}

func (s *FilterSettings) ID() string {
	return s.Name
}

func (s *FilterSettings) compile() (*filter, error) {
	return new(filter).compile(s)
}

func (s *FilterSettings) compileJoin() (*filterJoin, error) {
	return new(filterJoin).compile(s)
}
