package sdk

import (
	"fmt"
)

type IAggregationRule interface {
	ID() string
	Run()
	Feed(event *Event)
}

type AggregationRule struct {
	BaseRule
	aggregator       *simpleAggregator
	aggregationChain chan AggregationChainTask
}

func (ar *AggregationRule) Run() {
	ar.aggregator = newSimpleAggregator(
		ar.aggregation,
		ar.eventSpecs,
		ar.onTrigger,
	)
}

func (ar *AggregationRule) Feed(event *Event) {
	for _, eventSpec := range ar.eventSpecs {
		if eventSpec.filter.Pass([]*Event{event}) {
			ar.aggregator.feed(event, eventSpec)
			return
		}
	}
	ar.aggregationChain <- event
}

// This will be concurrently called by aggregation cells
func (ar *AggregationRule) onTrigger(trigger string, payload *triggerPayload) {
	if Debug {
		ar.printDebugInfo(trigger, payload)
	}

	if trigger == triggerEveryThreshold || trigger == triggerTimeout {
		payload.events[0].AggregationRuleName = ar.name
		ar.aggregationChain <- payload.events[0]
	}
}

func (ar *AggregationRule) printDebugInfo(trigger string, payload *triggerPayload) {
	eventSpecID := "-"

	if payload.eventSpec != nil {
		eventSpecID = payload.eventSpec.id
	}

	fmt.Printf(
		"Aggregation rule '%s' triggered by '%s' on '%s'\n",
		ar.name, eventSpecID, trigger,
	)
}
