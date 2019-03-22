package correlator

import (
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/correlation"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
)

type Processor struct {
	InputChannel     connectors.OutputChannel
	CorrelationRules []correlation.IRule
	Destinations     []destinations.IDestination
}

func (r *Processor) Start() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.worker()
	}
}

func (r *Processor) worker() {
	for input := range r.InputChannel {
		event := new(normalization.Event)
		if err := event.FromMsgPack(input.Data); err != nil {
			continue
		}

		for _, rule := range r.CorrelationRules {
			rule.Feed(event)
		}
	}
}

func (r *Processor) output(event *normalization.Event) {
	for _, dst := range r.Destinations {
		dst.Send(event)
	}
}
