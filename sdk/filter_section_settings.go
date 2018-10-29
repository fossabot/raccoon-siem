package sdk

type FilterSectionSettings struct {
	Or          bool     `yaml:"or,omitempty"`
	Not         bool     `yaml:"not,omitempty"`
	Expressions []string `yaml:"expressions,omitempty"`
}

func (s *FilterSectionSettings) compile() (*filterSection, error) {
	return new(filterSection).compile(s)
}

func (s *FilterSectionSettings) compileJoin() (*filterJoinSection, error) {
	return new(filterJoinSection).compile(s)
}
