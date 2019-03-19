package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
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

type Config struct {
	Kind             string                   `yaml:"kind,omitempty"`
	Name             string                   `yaml:"name,omitempty"`
	AggregationRules []aggregation.Config     `yaml:"aggregationRules,omitempty"`
	Filter           *filters.JoinConfig      `yaml:"filter,omitempty"`
	Triggers         map[string]TriggerConfig `yaml:"triggers,omitempty"`
	Window           time.Duration            `yaml:"window,omitempty"`
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
//	Values    map[string]string `yaml:"values,omitempty"`
//}
