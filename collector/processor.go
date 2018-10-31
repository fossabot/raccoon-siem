package collector

import (
	"github.com/tephrocactus/raccoon-siem/logger"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"log"
	"time"
)

type Processor struct {
	ParsingChannel     chan sdk.ProcessorTask
	AggregationChannel chan sdk.AggregationChainTask
	Workers            int
	Parsers            []sdk.IParser
	Filters            []sdk.IFilter
	AggregationRules   []sdk.IAggregationRule
	Sources            []sdk.ISource
	Destinations       []sdk.IDestination
	ID                 string
	Debug              bool
	hostname           string
	ip                 string
	logger             *logger.Instance
}

func (r *Processor) Start() *Processor {
	r.hostname = sdk.GetHostName()
	r.ip = sdk.GetIPAddress()

	for i := 0; i < r.Workers; i++ {
		go r.parsingRoutine()
		go r.aggregationRoutine()
	}

	sdk.RunAggregationRules(r.AggregationRules)

	if err := sdk.RunDestinations(r.Destinations); err != nil {
		log.Fatal(err)
	}

	if err := sdk.RunSources(r.Sources); err != nil {
		log.Fatal(err)
	}

	logLevel := logger.LevelError
	if r.Debug {
		logLevel = logger.LevelDebug
	}

	r.logger = logger.NewInstance(r.ID, r.Destinations, logLevel)
	return r
}

// Processes incoming events
func (r *Processor) parsingRoutine() {
channelLoop:
	for data := range r.ParsingChannel {
		if len(data) == 0 {
			continue
		}

		event, err := r.parse(data)

		if err != nil {
			if r.Debug {
				r.logger.Debug(err.Error(), &sdk.Event{
					Details: string(data),
				})
			}
			continue
		}

		for _, f := range r.Filters {
			if !f.Pass([]*sdk.Event{event}) {
				if r.Debug {
					r.logger.Debug("filtered out", &sdk.Event{
						Details: string(data),
					})
				}
				continue channelLoop
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
			continue
		}

		for _, ar := range r.AggregationRules {
			ar.Feed(event)
		}
	}
}

func (r *Processor) aggregationRoutine() {
	for event := range r.AggregationChannel {
		event.CollectorDNSName = r.hostname
		event.CollectorIPAddress = r.ip
		for _, dst := range r.Destinations {
			dst.Send(event)
		}
	}
}

// Parses events with parser/sub chain
func (r *Processor) parse(data []byte) (event *sdk.Event, err error) {
	for _, parser := range r.Parsers {
		event, err = parser.Parse(data, nil)

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
