package sdk

import "fmt"

const (
	defaultTTLSeconds = 2592000
)

type ActiveListServiceSettings struct {
	Name     string `yaml:"name,omitempty"`
	URL      string `yaml:"url,omitempty"`
	PoolSize int    `yaml:"poolSize,omitempty"`
}

func (s *ActiveListServiceSettings) ID() string {
	return s.Name
}

type ActiveListSettings struct {
	Name string `yaml:"name,omitempty"`
	TTL  int64  `yaml:"ttl,omitempty"`
}

func (s *ActiveListSettings) ID() string {
	return s.Name
}

func (s *ActiveListSettings) compile() (*activeListSpecification, error) {
	spec := new(activeListSpecification)

	if s.Name == "" {
		return nil, fmt.Errorf("active list must have a name")
	}

	spec.name = s.Name
	spec.ttl = s.TTL

	if spec.ttl == 0 {
		spec.ttl = defaultTTLSeconds
	}

	return spec, nil
}
