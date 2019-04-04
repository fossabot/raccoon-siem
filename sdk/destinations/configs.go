package destinations

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	DestinationNATS    = "nats"
	DestinationElastic = "elastic"
)

type InputChannel chan *normalization.Event

type Config struct {
	Id          string `json:"id,omitempty" db:"id,omitempty"`
	Name        string `json:"name,omitempty" db:"name,omitempty"`
	Kind        string `json:"kind,omitempty" db:"kind,omitempty"`
	URL         string `json:"url,omitempty" db:"url,omitempty"`
	Subject     string `json:"subject,omitempty" db:"subject"`
	Index       string `json:"index,omitempty" db:"index_name"`
	StaticIndex bool   `json:"staticIndex,omitempty" db:"static_index"`
	BatchSize   int    `json:"batchSize,omitempty" db:"batch_size"`
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("destination: name required")
	}

	switch r.Kind {
	case DestinationElastic:
		return r.validateElastic()
	case DestinationNATS:
		return r.validateNats()
	default:
		return fmt.Errorf("destination: unknown kind %s", r.Kind)
	}
}

func (r *Config) validateElastic() error {
	if r.URL == "" {
		return fmt.Errorf("destination: url required")
	}

	if r.Index == "" {
		return fmt.Errorf("destination: index required")
	}

	return nil
}

func (r *Config) validateNats() error {
	if r.URL == "" {
		return fmt.Errorf("destination: url required")
	}

	if r.Subject == "" {
		return fmt.Errorf("destination: subject required")
	}

	return nil
}

func (r *Config) ID() string {
	return r.Name
}
