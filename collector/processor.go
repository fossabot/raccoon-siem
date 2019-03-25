package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers"
	"time"
)

type Processor struct {
	hostname         string
	ipAddress        string
	metrics          *metrics
	inputChannel     connectors.OutputChannel
	normalizer       normalizers.INormalizer
	filters          []*filters.Filter
	enrichment       []enrichment.Config
	aggregationRules []*aggregation.Rule
	destinations     []destinations.IDestination
	workers          int
}

func (r *Processor) Start() {
	for i := 0; i < r.workers; i++ {
		go r.worker()
	}
}

func (r *Processor) worker() {
mainLoop:
	for input := range r.inputChannel {
		processingBegan := time.Now()
		r.metrics.eventReceived()

		event := r.normalizer.Normalize(input.Data, nil)
		if event == nil {
			r.metrics.eventProcessed()
			continue
		}

		for _, dropFilter := range r.filters {
			if dropFilter.Pass(event) {
				r.metrics.eventFiltered(dropFilter.ID())
				r.metrics.eventProcessed()
				continue mainLoop
			}
		}

		for _, config := range r.enrichment {
			enrichment.Enrich(config, event)
		}

		event.Timestamp = helpers.NowUnixMillis()
		event.ID = helpers.GetUUID()
		event.SourceID = input.Connector
		event.CollectorDNSName = r.hostname
		event.CollectorIPAddress = r.ipAddress

		for _, rule := range r.aggregationRules {
			if rule.Feed(event) {
				r.metrics.eventAggregated(rule.ID())
				r.metrics.eventProcessed()
				r.metrics.processingTook(time.Since(processingBegan))
				continue mainLoop
			}
		}

		r.output(event)
		r.metrics.eventProcessed()
		r.metrics.processingTook(time.Since(processingBegan))
	}
}

func (r *Processor) output(event *normalization.Event) {
	for _, dst := range r.destinations {
		dst.Send(event)
		r.metrics.eventSent(dst.ID())
	}
}
