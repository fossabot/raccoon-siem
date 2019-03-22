package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
	"runtime"
)

type Processor struct {
	InputChannel     connectors.OutputChannel
	Normalizer       normalizers.INormalizer
	DropFilters      []*filters.Filter
	EnrichConfigs    []enrichment.Config
	AggregationRules []aggregation.Rule
	Destinations     []destinations.IDestination
}

func (r *Processor) Start() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.worker()
	}
}

func (r *Processor) worker() {
mainLoop:
	for input := range r.InputChannel {
		event := r.Normalizer.Normalize(input.Data, nil)
		if event == nil {
			continue
		}

		for _, dropFilter := range r.DropFilters {
			if dropFilter.Pass(event) {
				continue mainLoop
			}
		}

		for _, config := range r.EnrichConfigs {
			enrichment.Enrich(config, event)
		}

		event.Timestamp = helpers.NowUnixMillis()
		event.ID = helpers.GetUUID()
		event.SourceID = input.Connector

		for _, rule := range r.AggregationRules {
			if rule.Feed(event) {
				continue mainLoop
			}
		}

		r.output(event)
	}
}

func (r *Processor) output(event *normalization.Event) {
	for _, dst := range r.Destinations {
		dst.Send(event)
	}
}
