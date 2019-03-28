package correlation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

const (
	TriggerFirstEvent           = "firstEvent"
	TriggerSubsequentEvents     = "subsequentEvents"
	TriggerEveryEvent           = "everyEvent"
	TriggerFirstThreshold       = "firstThreshold"
	TriggerSubsequentThresholds = "subsequentThresholds"
	TriggerEveryThreshold       = "everyThreshold"
	TriggerTimeout              = "timeout"
)

const (
	RuleKindCommon   = "common"
	RuleKindRecovery = "recovery"
)

const (
	BusSubject = "raccoon-correlation"
)

type OutputFn func(event *normalization.Event)

type Config struct {
	Name            string                   `json:"name,omitempty"`
	Kind            string                   `json:"kind,omitempty"`
	Selectors       []EventSelector          `json:"selectors,omitempty"`
	IdenticalFields []string                 `json:"identicalFields,omitempty"`
	UniqueFields    []string                 `json:"uniqueFields,omitempty"`
	Filter          *filters.JoinConfig      `json:"filter,omitempty"`
	Triggers        map[string]TriggerConfig `json:"triggers,omitempty"`
	Window          int64                    `json:"window,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("correlation rule: name required")
	}

	if r.Filter != nil {
		if err := r.Filter.Validate(); err != nil {
			return err
		}
	}

	if r.Triggers == nil {
		return fmt.Errorf("correlation rule: triggers required")
	}

	for k := range r.Triggers {
		switch k {
		case TriggerFirstEvent, TriggerSubsequentEvents, TriggerEveryEvent,
			TriggerFirstThreshold, TriggerSubsequentThresholds, TriggerEveryThreshold, TriggerTimeout:
		default:
			return fmt.Errorf("correlation rule: unknown trigger %s", k)
		}

		t := r.Triggers[k]
		if err := t.Validate(); err != nil {
			return err
		}
	}

	if len(r.Selectors) == 0 {
		return fmt.Errorf("correlation rule: event selectors required")
	}

	for i := range r.Selectors {
		if err := r.Selectors[i].Validate(); err != nil {
			return err
		}
	}

	if r.Kind == RuleKindRecovery {
		if err := r.validateRecovery(); err != nil {
			return err
		}
	}

	for _, f := range r.IdenticalFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("aggregation rule: invalid event field %s", f)
		}
	}

	for _, f := range r.UniqueFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("aggregation rule: invalid event field %s", f)
		}
	}

	return nil
}

func (r *Config) validateRecovery() error {
	hasRecoverySelector := false
	for _, selector := range r.Selectors {
		if selector.Recovery {
			hasRecoverySelector = true
			break
		}
	}

	if !hasRecoverySelector {
		return fmt.Errorf("correlation rule: at least one recovery event selector required")
	}

	return nil
}

func (r *Config) WindowDuration() time.Duration {
	return time.Duration(r.Window) * time.Second
}

type EventSelector struct {
	Tag       string         `json:"tag,omitempty"`
	Filter    filters.Config `json:"filter,omitempty"`
	Threshold int            `json:"threshold,omitempty"`
	Recovery  bool           `json:"recovery,omitempty"`
}

func (r *EventSelector) Validate() error {
	if r.Tag == "" {
		return fmt.Errorf("event selector: tag required")
	}

	if err := r.Filter.Validate(); err != nil {
		return err
	}

	return nil
}

type TriggerConfig struct {
	Actions []ActionConfig `json:"actions,omitempty"`
}

func (r *TriggerConfig) Validate() error {
	if len(r.Actions) == 0 {
		return fmt.Errorf("trigger: actions required")
	}

	for i := range r.Actions {
		if err := r.Actions[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ActionConfig struct {
	Kind       string                   `json:"kind,omitempty"`
	Release    actions.ReleaseConfig    `json:"release,omitempty"`
	ActiveList actions.ActiveListConfig `json:"activeList,omitempty"`
}

func (r *ActionConfig) Validate() error {
	switch r.Kind {
	case actions.KindRelease:
		return r.Release.Validate()
	case actions.KindActiveList:
		return r.ActiveList.Validate()
	default:
		return fmt.Errorf("action: unknown kind %s", r.Kind)
	}
}
