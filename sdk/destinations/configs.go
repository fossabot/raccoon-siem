package destinations

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

const (
	DestinationConsole = "console"
	DestinationNATS    = "nats"
	DestinationElastic = "elastic"
)

type InputChannel chan *normalization.Event

type Config struct {
	Name        string `yaml:"name,omitempty"`
	Kind        string `yaml:"kind,omitempty"`
	URL         string `yaml:"url,omitempty"`
	Subject     string `yaml:"subject,omitempty"`
	Index       string `yaml:"index,omitempty"`
	StaticIndex bool   `yaml:"staticIndex,omitempty"`
	BatchSize   int    `yaml:"batchSize,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}
