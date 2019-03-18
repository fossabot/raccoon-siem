package correlator

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type Processor struct {
	CorrelationChannel      chan *sdk.ProcessorTask
	CorrelationChainChannel chan sdk.CorrelationChainTask
	CorrelationRules        []sdk.ICorrelationRule
	Workers                 int
	Debug                   bool
	hostname                string
	ip                      string
}

func (r *Processor) Start() error {
	r.hostname = sdk.GetHostName()
	r.ip = sdk.GetIPAddress()

	for i := 0; i < r.Workers; i++ {
		go r.correlationRoutine()
		go r.correlationChainRoutine()
	}

	sdk.RunCorrelationRules(r.CorrelationRules)

	return nil
}

// Processes incoming events
func (r *Processor) correlationRoutine() {
	for input := range r.CorrelationChannel {
		event := new(normalization.Event)
		if err := event.FromMsgPack(input.Data); err != nil {
			continue
		}

		event.CorrelatorDNSName = r.hostname
		event.CorrelatorIPAddress = r.ip

		for _, rule := range r.CorrelationRules {
			rule.Feed(event)
		}
	}
}

// Processes correlated events
func (r *Processor) correlationChainRoutine() {
	for event := range r.CorrelationChainChannel {
		event.CorrelatorDNSName = r.hostname
		event.CorrelatorIPAddress = r.ip
		fmt.Println(event)
	}
}
