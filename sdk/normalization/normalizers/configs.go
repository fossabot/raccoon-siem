package normalizers

type Config struct {
	Name          string          `yaml:"name,omitempty"`
	Kind          string          `yaml:"kind,omitempty"`
	Expressions   []string        `yaml:"expressions,omitempty"`
	PairDelimiter string          `yaml:"pairDelimiter,omitempty"`
	KVDelimiter   string          `yaml:"kvDelimiter,omitempty"`
	Mapping       []MappingConfig `yaml:"mapping,omitempty"`
}

func (s *Config) ID() string {
	return s.Name
}

type ExtraConfig struct {
	Normalizer   Config      `yaml:"normalizer"`
	TriggerField string      `yaml:"triggerField,omitempty"`
	TriggerValue string      `yaml:"triggerValue,omitempty"`
	triggerValue []byte      `json:"-" yaml:"-"`
	normalizer   INormalizer `json:"-" yaml:"-"`
}

type MappingConfig struct {
	SourceField string        `yaml:"sourceField,omitempty"`
	EventField  string        `yaml:"eventField,omitempty"`
	Extra       []ExtraConfig `yaml:"extra,omitempty"`
}
