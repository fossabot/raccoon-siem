package connectors

import (
	"fmt"
)

type Config struct {
	Name       string `json:"name,omitempty"`
	Kind       string `json:"kind,omitempty"`
	URL        string `json:"url"`
	Proto      string `json:"proto"`
	Subject    string `json:"subject"`
	Queue      string `json:"queue"`
	Delimiter  string `json:"delimiter"`
	BufferSize int    `json:"bufferSize"`
	MaxLen     int    `json:"maxLen"`
	Partition  int    `json:"partition"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("connector: name required")
	}

	switch r.Kind {
	case connectorListener:
		return r.validateListener()
	case connectorNetflow:
		return r.validateNetflow()
	case connectorNats:
		return r.validateNats()
	default:
		return fmt.Errorf("connector: unknown kind %s", r.Kind)
	}
}

func (r *Config) validateListener() error {
	if r.URL == "" {
		return fmt.Errorf("connector: url required")
	}

	switch r.Proto {
	case "tcp", "udp":
	default:
		return fmt.Errorf("connector: protocol must be tcp or udp")
	}

	if r.Delimiter != "" && len([]byte(r.Delimiter)) != 1 {
		return fmt.Errorf("connector: delimiter must be single byte ASCII character")
	}

	return nil
}

func (r *Config) validateNetflow() error {
	if r.URL == "" {
		return fmt.Errorf("connector: url required")
	}
	return nil
}

func (r *Config) validateNats() error {
	if r.URL == "" {
		return fmt.Errorf("connector: url required")
	}

	if r.Subject == "" {
		return fmt.Errorf("connector: subject required")
	}

	return nil
}
