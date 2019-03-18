package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
	"runtime"
	"time"
)

type Processor struct {
	InputChannel       connectors.OutputChannel
	AggregationChannel chan normalization.Event
	OutputChannel      chan normalization.Event
	Normalizer         normalizers.INormalizer
	DropFilters        []*filters.Filter
	AggregationRules   []aggregation.Rule
	EnrichConfigs      []enrichment.EnrichConfig
}

func (r *Processor) Start() error {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.worker()
	}
	return nil
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
		
		event.Timestamp = time.Now()
		event.ID = sdk.GetUUID()
		event.SourceID = input.Connector

		for _, rule := range r.AggregationRules {
			if rule.Feed(event) {
				continue mainLoop
			}
		}

		r.OutputChannel <- *event
	}
}
