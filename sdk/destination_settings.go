package sdk

type DestinationSettings struct {
	Name                string `yaml:"name,omitempty"`
	Kind                string `yaml:"kind,omitempty"`
	URL                 string `yaml:"url,omitempty"`
	Channel             string `yaml:"channel,omitempty"`
	Index               string `yaml:"index,omitempty"`
	RotateIndex         bool   `yaml:"rotateIndex,omitempty"`
	Size                int    `yaml:"size,omitempty"`
	BenchmarkEventCount int    `yaml:"benchmarkEventCount,omitempty"`
}

func (s *DestinationSettings) ID() string {
	return s.Name
}
