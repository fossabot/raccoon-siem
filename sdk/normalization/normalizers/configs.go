package normalizers

import (
	"errors"
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"strings"
)


type Config struct {
	Name          string          `json:"name,omitempty"`
	Kind          string          `json:"kind,omitempty"`
	Expressions   []string        `json:"expressions,omitempty"`
	PairDelimiter string          `json:"pairDelimiter,omitempty"`
	KVDelimiter   string          `json:"kvDelimiter,omitempty"`
	Mapping       []MappingConfig `json:"mapping,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return errors.New("normalizer: name required")
	}

	switch r.Kind {
	case KindJSON, KindSyslog, KindCEF:
	case KindRegexp:
		if len(r.Expressions) == 0 {
			return errors.New("normalizer: expressions required")
		}
	case KindKV:
		if r.KVDelimiter == "" || r.PairDelimiter == "" {
			return errors.New("normalizer: delimiters required")
		}
	default:
		return fmt.Errorf("normalizer: unknown kind %s", r.Kind)
	}

	if len(r.Mapping) == 0 {
		return errors.New("normalizer: mapping required")
	}

	for i := range r.Mapping {
		if err := r.Mapping[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type MappingConfig struct {
	SourceField string        `json:"sourceField,omitempty"`
	EventField  string        `json:"eventField,omitempty"`
	Extra       []ExtraConfig `json:"extra,omitempty"`
	Label       string        `json:"label,omitempty"`
	labelField  string        `json:"-"`
}

func (r *MappingConfig) Validate() error {
	if r.SourceField == "" {
		return errors.New("mapping: source field required")
	}

	if !helpers.EventFieldHasGetter(r.EventField) {
		return fmt.Errorf("mapping: invalid event field %s", r.EventField)
	}

	if r.Label != "" {
		if strings.Index(r.EventField, "User") != 0 {
			return fmt.Errorf("only User* event fields may have a label")
		}
		r.labelField = r.EventField + "Label"
	}

	for i := range r.Extra {
		if err := r.Extra[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ExtraConfig struct {
	Normalizer   Config      `json:"normalizer"`
	TriggerField string      `json:"triggerField,omitempty"`
	TriggerValue string      `json:"triggerValue,omitempty"`
	triggerValue []byte      `json:"-"`
	normalizer   INormalizer `json:"-"`
}

func (r *ExtraConfig) Validate() error {
	if r.TriggerValue != "" && r.TriggerField == "" {
		return errors.New("extra: trigger field required")
	}
	return r.Normalizer.Validate()
}

type extraMarker struct {
	extPtr *ExtraConfig
	value  []byte
}
