package correlator

import (
	"errors"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

type Processor struct {
	CorrelationChannel      chan *sdk.ProcessorTask
	CorrelationChainChannel chan sdk.CorrelationChainTask
	Workers                 int
	Parsers                 []sdk.IParser
	CorrelationRules        []sdk.ICorrelationRule
	Connectors              []sdk.IConnector
	Destinations            []sdk.IDestination
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

	if err := sdk.RunDestinations(r.Destinations); err != nil {
		return err
	}

	if err := sdk.RunConnectors(r.Connectors); err != nil {
		return err
	}

	return nil
}

// Processes incoming events
func (r *Processor) correlationRoutine() {
	for task := range r.CorrelationChannel {
		event, err := r.parse(task.Data)

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
