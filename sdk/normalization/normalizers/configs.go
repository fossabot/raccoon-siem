package normalizers

import (
	"encoding/json"
	"errors"
)

type Config struct {
	Name          string          `yaml:"name,omitempty" json:"name"`
	Kind          string          `yaml:"kind,omitempty" json:"kind"`
	Expressions   []string        `yaml:"expressions,omitempty" json:"expressions"`
	PairDelimiter string          `yaml:"pairDelimiter,omitempty" json:"pairDelimiter"`
	KVDelimiter   string          `yaml:"kvDelimiter,omitempty" json:"kvDelimiter"`
	Mapping       []MappingConfig `yaml:"mapping,omitempty" json:"mapping"`
}

func (r *Config) ID() string {
	return r.Name
}

type ExtraConfig struct {
	Normalizer   Config      `yaml:"normalizer" json:"normalizer"`
	TriggerField string      `yaml:"triggerField,omitempty" json:"triggerField"`
	TriggerValue string      `yaml:"triggerValue,omitempty" json:"triggerValue"`
	triggerValue []byte      `json:"-" yaml:"-"`
	normalizer   INormalizer `json:"-" yaml:"-"`
}

type MappingConfig struct {
	SourceField string        `yaml:"sourceField,omitempty" json:"sourceField"`
	EventField  string        `yaml:"eventField,omitempty" json:"eventField"`
	Extra       []ExtraConfig `yaml:"extra,omitempty" json:"extra"`
}

func (r *Config) Unmarshal(data []byte) error {
	config := Config{}
	err := json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	err = config.validate()

	if err == nil {
		*r = config
	}
	return err
}

func (r *Config) validate() error {
	if r.Name == "" {
		return errors.New("name required")
	}

	if r.Kind == "" {
		return errors.New("kind required")
	}

	if r.Kind == KindRegexp {
		if len(r.Expressions) == 0 {
			return errors.New("expressions required")
		}
	} else if r.Kind == KindKV {
		if r.KVDelimiter == "" || r.PairDelimiter == "" {
			return errors.New("delimiters required")
		}
	}

	if len(r.Mapping) == 0 {
		return errors.New("mapping required")
	}

	for i := range r.Mapping {
		if r.Mapping[i].SourceField == "" {
			return errors.New("source field required")
		}

		if r.Mapping[i].EventField == "" {
			return errors.New("event field required")
		}

		for j := range r.Mapping[i].Extra {
			extraConfig := r.Mapping[i].Extra[j]
			if extraConfig.TriggerValue != "" && extraConfig.TriggerField == "" {
				return errors.New("trigger field required")
			}

			err := extraConfig.Normalizer.validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}