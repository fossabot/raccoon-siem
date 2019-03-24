package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

const (
	TriggerFirstEvent           = "firstEvent"
	TriggerSubsequentEvents     = "subsequentEvents"
	TriggerEveryEvent           = "everyEvent"
	TriggerFirstThreshold       = "firstTrigger"
	TriggerSubsequentThresholds = "subsequentTriggers"
	TriggerEveryThreshold       = "everyTrigger"
	TriggerTimeout              = "timeout"
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
	Kind       string                   `yaml:"kind,omitempty"`
	Release    actions.ReleaseConfig    `yaml:"release,omitempty"`
	ActiveList actions.ActiveListConfig `yaml:"activeList,omitempty"`
}
