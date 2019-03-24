package correlator

import (
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/correlation"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
	"time"
)

type Processor struct {
	hostname         string
	ipAddress        string
	metrics          *metrics
	inputChannel     connectors.OutputChannel
	correlationRules []correlation.IRule
	destinations     []destinations.IDestination
	workers          int
}

func (r *Processor) Start() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.worker()
	}
}

func (r *Processor) worker() {
	for input := range r.inputChannel {
		processingBegan := time.Now()
		r.metrics.eventReceived()

		event := new(normalization.Event)
		if err := event.FromMsgPack(input.Data); err != nil {
			r.metrics.eventProcessed()
			continue
		}

		for _, rule := range r.correlationRules {
			rule.Feed(event)
		}

		r.metrics.eventProcessed()
		r.metrics.processingTook(time.Since(processingBegan))
	}
}

func (r *Processor) output(event *normalization.Event) {
	r.metrics.eventCorrelated(event.CorrelationRuleName)
	for _, dst := range r.destinations {
		dst.Send(event)
		r.metrics.eventSent(dst.ID())
	}
}
