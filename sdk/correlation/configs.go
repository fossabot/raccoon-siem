package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
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
	ActionRelease = "release"
	ActionMutate  = "mutate"
	ActionAL      = "al"
)

const (
	ValueSourceKindConst = "const"
	ValueSourceKindEvent = "event"
	ValueSourceKindAL    = "al"
)

type Config struct {
	Name             string                   `yaml:"name,omitempty"`
	AggregationRules []aggregation.Config     `yaml:"aggregationRules,omitempty"`
	Filter           *filters.JoinConfig      `yaml:"filter,omitempty"`
	Triggers         map[string]TriggerConfig `yaml:"triggers,omitempty"`
	Unexpected       []string                 `yaml:"unexpected,omitempty"`
	Window           time.Duration            `yaml:"window,omitempty"`
}

type TriggerConfig struct {
	Kind    string         `yaml:"kind,omitempty"`
	Actions []ActionConfig `yaml:"actions,omitempty"`
}

type ActionConfig struct {
	Kind   string         `yaml:"kind,omitempty"`
	Mutate []MutateConfig `yaml:"mutate,omitempty"`
	//ActiveList []ActiveListActionConfig `yaml:"activeList,omitempty"`
}

type MutateConfig struct {
	Field            string      `yaml:"field,omitempty"`
	Constant         interface{} `yaml:"constant,omitempty"`
	KeyFields        []string    `yaml:"keyFields,omitempty"`
	ValueSourceKind  string      `yaml:"valueSourceKind,omitempty"`
	ValueSourceName  string      `yaml:"valueSourceName,omitempty"`
	ValueSourceField string      `yaml:"valueSourceField,omitempty"`
}

//type ActiveListActionConfig struct {
//	Name      string            `yaml:"name,omitempty"`
//	KeyFields []string          `yaml:"keyFields,omitempty"`
//	Op        string            `yaml:"op,omitempty"`
//	Values    map[string]string `yaml:"values,omitempty"`
//}
