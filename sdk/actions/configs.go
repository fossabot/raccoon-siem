package actions

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

const (
	KindRelease    = "release"
	KindActiveList = "activeList"
)

type ReleaseConfig struct {
	EnrichmentConfigs []enrichment.Config `json:"enrichment,omitempty"`
}

func (r *ReleaseConfig) Validate() error {
	if len(r.EnrichmentConfigs) == 0 {
		return fmt.Errorf("release action: enrichment required")
	}

	for i := range r.EnrichmentConfigs {
		if err := r.EnrichmentConfigs[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ActiveListConfig struct {
	Name      string                `json:"name,omitempty"`
	Op        string                `json:"op,omitempty"`
	KeyFields []string              `json:"keyFields,omitempty"`
	Mapping   []activeLists.Mapping `json:"mapping,omitempty"`
	EventTag  string                `json:"eventTag,omitempty"`
}

func (r *ActiveListConfig) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("active list action: name required")
	}

	switch r.Op {
	case activeLists.OpSet, activeLists.OpDel, activeLists.OpGet:
	default:
		return fmt.Errorf("active list action: unknown operation %s", r.Op)
	}

	if len(r.KeyFields) == 0 {
		return fmt.Errorf("active list action: key fields required")
	}

	for _, f := range r.KeyFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("active list action: invalid event field %s", f)
		}
	}

	if len(r.Mapping) == 0 {
		return fmt.Errorf("active list action: mapping required")
	}

	for i := range r.Mapping {
		if err := r.Mapping[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}
