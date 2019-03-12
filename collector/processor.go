package collector

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/logger"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"time"
)

type Processor struct {
	ParsingChannel     chan *sdk.ProcessorTask
	AggregationChannel chan sdk.AggregationChainTask
	Workers            int
	Parsers            []sdk.IParser
	Filters            []sdk.IFilter
	AggregationRules   []sdk.IAggregationRule
	Connectors         []sdk.IConnector
	Destinations       []sdk.IDestination
	ID                 string
	MetricsPort        string
	Debug              bool
	PrintRaw           bool
	hostname           string
	ip                 string
	logger             *logger.Instance
	metrics            *metrics
}

func (r *Processor) Start() error {
	r.hostname = sdk.GetHostName()
	r.ip = sdk.GetIPAddress()

	for i := 0; i < r.Workers; i++ {
		go r.parsingRoutine()
		go r.aggregationRoutine()
	}

	sdk.RunAggregationRules(r.AggregationRules)

	if err := sdk.RunDestinations(r.Destinations); err != nil {
		return err
	}

	logLevel := logger.LevelError
	if r.Debug {
		logLevel = logger.LevelDebug
	}

	r.logger = logger.NewInstance(r.ID, r.Destinations, logLevel)
	r.metrics = newMetrics(r.MetricsPort).runServer()

	if err := sdk.RunConnectors(r.Connectors); err != nil {
		return err
	}

	return nil
}

// Processes incoming events
func (r *Processor) parsingRoutine() {
	for task := range r.ParsingChannel {
		r.metrics.registerEventInput(task.Connector)
		if len(task.Data) != 0 {
			start := time.Now()

			if r.PrintRaw {
				fmt.Println(string(task.Data))
			}

			r.processEvent(task)
			r.metrics.registerOverallProcessingDuration(start)
		}
	}
}

func (r *Processor) processEvent(task *sdk.ProcessorTask) {
	event, err := r.parse(task.Data)

	if err != nil {
		if r.Debug {
			r.logger.Debug(err.Error(), &sdk.Event{Details: string(task.Data)})
		}
		return
	}

	event.SourceID = task.Connector

	for _, f := range r.Filters {
		if !f.Pass([]*sdk.Event{event}) {
			r.metrics.registerEventFiltration(f.ID(), task.Connector)
			if r.Debug {
				r.logger.Debug("filtered out", &sdk.Event{Details: string(task.Data)})
			}
			return
		}
	}

	if event.OriginTimestamp.IsZero() {
		event.OriginTimestamp = time.Now()
	}

	event.ID = sdk.GetUUID()
	event.StartTime = event.OriginTimestamp
	event.EndTime = event.OriginTimestamp

	if len(r.AggregationRules) == 0 {
		r.AggregationChannel <- event
		return
	}

	for _, ar := range r.AggregationRules {
		ar.Feed(event)
	}
}

func (r *Processor) aggregationRoutine() {
	for event := range r.AggregationChannel {
		event.CollectorDNSName = r.hostname
		event.CollectorIPAddress = r.ip

		if event.AggregatedEventCount > 0 {
			r.metrics.registerEventAggregation(event.AggregationRuleName, event.AggregatedEventCount, event.SourceID)
		}

		for _, dst := range r.Destinations {
			start := time.Now()
			dst.Send(event)
			r.metrics.registerSerializationDuration(dst.ID(), start)
			r.metrics.registerEventOutput(event.SourceID)
		}
	}
}

// Parses events with parser/sub chain
func (r *Processor) parse(data []byte) (event *sdk.Event, err error) {
	for _, parser := range r.Parsers {
		start := time.Now()
		event, err = parser.Parse(data, nil)
		r.metrics.registerParsingDuration(parser.ID(), start)

		if err != nil {
			continue
		}

		break
	}

	if err != nil {
		err = sdk.ErrAllParsersFailed
	}

	return
}
