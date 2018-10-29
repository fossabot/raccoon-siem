package sdk

type ParserVariantSettings struct {
	Regexp  string   `yaml:"regexp,omitempty"`
	Mapping []string `yaml:"mapping,omitempty"`
}

type ParserSettings struct {
	Name            string                  `yaml:"name,omitempty"`
	Kind            string                  `yaml:"kind,omitempty"`
	Regexp          string                  `yaml:"regexp,omitempty"`
	Subs            []string                `yaml:"subs,omitempty"`
	Mapping         []string                `yaml:"mapping,omitempty"`
	Variables       []string                `yaml:"vars,omitempty"`
	Rewrites        []string                `yaml:"rewrites,omitempty"`
	Variants        []ParserVariantSettings `yaml:"variants,omitempty"`
	KVDelimiter     string                  `yaml:"kvDelimiter,omitempty"`
	KVPairDelimiter string                  `yaml:"kvPairDelimiter,omitempty"`
	Root            bool                    `yaml:"root,omitempty"`
}

func (s *ParserSettings) ID() string {
	return s.Name
}

func (s *ParserSettings) Compile() (*parserSpecification, error) {
	return new(parserSpecification).compile(s)
}
