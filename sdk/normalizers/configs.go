package normalizers

type Config struct {
	Name          string          `yaml:"name,omitempty"`
	Kind          string          `yaml:"kind,omitempty"`
	Expressions   []string        `yaml:"expressions,omitempty"`
	PairDelimiter byte            `yaml:"pair_delimiter,omitempty"`
	KVDelimiter   byte            `yaml:"kv_delimiter,omitempty"`
	Mapping       []MappingConfig `yaml:"mapping,omitempty"`
}

func (s *Config) ID() string {
	return s.Name
}

type ExtraConfig struct {
	NormalizerName string      `yaml:"normalizer_name,omitempty"`
	TriggerField   string      `yaml:"trigger_field,omitempty"`
	TriggerValue   []byte      `yaml:"trigger_value,omitempty"`
	Normalizer     INormalizer `yaml:"-"`
}

type MappingConfig struct {
	SourceField string       `yaml:"source_field,omitempty"`
	EventField  string       `yaml:"event_field,omitempty"`
	TimeFormat  string       `yaml:"time_format,omitempty"`
	Extra       *ExtraConfig `yaml:"extra,omitempty"`
	timeFormat  byte
}
