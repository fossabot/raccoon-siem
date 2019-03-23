package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

const (
	TriggerFirstEvent           = "fe"
	TriggerSubsequentEvents     = "se"
	TriggerEveryEvent           = "ee"
	TriggerFirstThreshold       = "ft"
	TriggerSubsequentThresholds = "st"
	TriggerEveryThreshold       = "et"
	TriggerTimeout              = "to"
)

const (
	RuleKindCommon   = "common"
	RuleKindRecovery = "recover"
)

const (
	BusSubject = "raccoon-correlation"
)

type OutputFn func(event *normalization.Event)

type Config struct {
	Kind            string                   `yaml:"kind,omitempty"`
	Name            string                   `yaml:"name,omitempty"`
	Selectors       []EventSelector          `yaml:"selectors,omitempty"`
	IdenticalFields []string                 `yaml:"identicalFields,omitempty"`
	UniqueFields    []string                 `yaml:"uniqueFields,omitempty"`
	Filter          *filters.JoinConfig      `yaml:"filter,omitempty"`
	Triggers        map[string]TriggerConfig `yaml:"triggers,omitempty"`
	Window          time.Duration            `yaml:"window,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}

type EventSelector struct {
	Tag       string         `yaml:"tag,omitempty"`
	Filter    filters.Config `yaml:"filter,omitempty"`
	Threshold int            `yaml:"threshold,omitempty"`
	Recovery  bool           `yaml:"recovery,omitempty"`
}

type TriggerConfig struct {
	Actions []ActionConfig `yaml:"actions,omitempty"`
}

type ActionConfig struct {
	Kind    string                `yaml:"kind,omitempty"`
	Release actions.ReleaseConfig `yaml:"release,omitempty"`
	//ActiveList []ActiveListActionConfig `yaml:"activeList,omitempty"`
}

//type ActiveListActionConfig struct {
//	Name      string            `yaml:"name,omitempty"`
//	KeyFields []string          `yaml:"keyFields,omitempty"`
//	Op        string            `yaml:"op,omitempty"`
//	Record    map[string]string `yaml:"values,omitempty"`
//}
