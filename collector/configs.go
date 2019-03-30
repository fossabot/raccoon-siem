package collector

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers"
)

type Config struct {
	Name         string                `json:"name,omitempty"`
	Connector    connectors.Config     `json:"connector,omitempty"`
	Normalizer   normalizers.Config    `json:"normalizer,omitempty"`
	Filters      []filters.Config      `json:"filters,omitempty"`
	Enrichment   []enrichment.Config   `json:"enrichment,omitempty"`
	Rules        []aggregation.Config  `json:"rules,omitempty"`
	Destinations []destinations.Config `json:"destinations,omitempty"`
	ActiveLists  []activeLists.Config  `json:"activeLists,omitempty"`
	Dictionaries []dictionaries.Config `json:"dictionaries,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("collector: name required")
	}

	if err := r.Connector.Validate(); err != nil {
		return err
	}

	if err := r.Normalizer.Validate(); err != nil {
		return err
	}

	for i := range r.Filters {
		if err := r.Filters[i].Validate(); err != nil {
			return err
		}
	}

	for i := range r.Enrichment {
		if err := r.Enrichment[i].Validate(); err != nil {
			return err
		}
	}

	for i := range r.Rules {
		if err := r.Rules[i].Validate(); err != nil {
			return err
		}
	}

	for i := range r.ActiveLists {
		if err := r.ActiveLists[i].Validate(); err != nil {
			return err
		}
	}

	for i := range r.Dictionaries {
		if err := r.Dictionaries[i].Validate(); err != nil {
			return err
		}
	}

	if len(r.Destinations) == 0 {
		return fmt.Errorf("collector: destinations required")
	}

	for i := range r.Destinations {
		if err := r.Destinations[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type cmdFlags struct {
	ID          string
	ConfigPath  string
	CoreURL     string
	BusURL      string
	StorageURL  string
	MetricsPort string
	Workers     int
	Debug       bool
}
