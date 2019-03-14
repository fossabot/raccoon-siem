package connectors

type Config struct {
	Name       string `yaml:"name,omitempty"`
	Kind       string `yaml:"kind,omitempty"`
	URL        string `yaml:"url"`
	Protocol   string `yaml:"protocol"`
	Subject    string `yaml:"subject"`
	Queue      string `yaml:"queue"`
	Delimiter  byte   `yaml:"delimiter"`
	BufferSize int    `yaml:"buffer_size"`
	MaxLen     int    `yaml:"max_len"`
}

func (s *Config) ID() string {
	return s.Name
}
