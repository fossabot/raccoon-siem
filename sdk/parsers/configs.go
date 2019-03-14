package parsers

type BaseConfig struct {
	Name string
}

type ParserVariantSettings struct {
	Regexp  string   `yaml:"regexp,omitempty"`
	Mapping []string `yaml:"mapping,omitempty"`
}

type UserConfig struct {
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

func (s *UserConfig) ID() string {
	return s.Name
}
