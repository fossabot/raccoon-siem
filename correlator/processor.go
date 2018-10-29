package correlator

import (
	"errors"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

type Processor struct {
	CorrelationChannel      chan sdk.ProcessorTask
	CorrelationChainChannel chan sdk.CorrelationChainTask
	Workers                 int
	Parsers                 []sdk.IParser
	CorrelationRules        []sdk.ICorrelationRule
	Sources                 []sdk.ISource
	Destinations            []sdk.IDestination
	Debug                   bool
	hostname                string
	ip                      string
}

func (r *Processor) Start() *Processor {
	r.hostname = sdk.GetHostName()
	r.ip = sdk.GetIPAddress()

	for i := 0; i < r.Workers; i++ {
		go r.correlationRoutine()
		go r.correlationChainRoutine()
	}

	sdk.RunCorrelationRules(r.CorrelationRules)

	err := sdk.RunDestinations(r.Destinations)
	sdk.PanicOnError(err)

	err = sdk.RunSources(r.Sources)
	sdk.PanicOnError(err)

	return r
}

// Processes incoming events
func (r *Processor) correlationRoutine() {
	for data := range r.CorrelationChannel {
		event, err := r.parse(data)

		if err != nil {
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

		for _, dst := range r.Destinations {
			dst.Send(event)
		}
	}
}

// Parses incoming events
func (r *Processor) parse(data []byte) (event *sdk.Event, err error) {
	for _, parser := range r.Parsers {
		event, err = parser.Parse(data, nil)

		if err != nil {
			continue
		}

		break
	}

	if err != nil {
		err = errors.New("all parsers failed")
	}

	return
}
