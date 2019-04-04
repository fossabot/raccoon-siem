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
	"log"
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
	debug            bool
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
			continue
		}

		for _, filter := range r.filters {
			if !filter.Pass(event) {
				r.metrics.eventFiltered(filter.ID())
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
				r.metrics.processingTook(time.Since(processingBegan))
				continue mainLoop
			}
		}

		r.metrics.processingTook(time.Since(processingBegan))

		outputBegan := time.Now()
		r.output(event)
		r.metrics.outputTook(time.Since(outputBegan))
	}
}

func (r *Processor) output(event *normalization.Event) {
	if r.debug {
		log.Println(event)
	}

	if encodedEvent, err := event.ToJSON(); err == nil {
		for _, dst := range r.destinations {
			dst.Send(encodedEvent)
			r.metrics.eventSent(dst.ID())
		}
	}
}
