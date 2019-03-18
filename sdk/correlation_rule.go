package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

type ICorrelationRule interface {
	ID() string
	Run()
	Feed(event *normalization.Event)
}

type CorrelationRule struct {
	BaseRule
	aggregator       *aggregator
	filter           filters.IFilter
	actions          actionsByTrigger
	correlationChain chan CorrelationChainTask
}

func (cr *CorrelationRule) Run() {
	cr.aggregator = newAggregator(
		cr.aggregation,
		cr.eventSpecs,
		cr.onTrigger,
		cr.onRootTrigger,
	)
}

func (cr *CorrelationRule) Feed(event *normalization.Event) {
	for _, eventSpec := range cr.eventSpecs {
		if eventSpec.filter.Pass(event) {
			cr.aggregator.feed(event, eventSpec)
			break
		}
	}
}

// This will be concurrently called by aggregation cells
func (cr *CorrelationRule) onTrigger(trigger string, payload *triggerPayload) {
	if Debug {
		eventSpecID := "-"
		if payload.eventSpec != nil {
			eventSpecID = payload.eventSpec.id
		}
		fmt.Printf(
			"Rule '%s' triggered by '%s' on '%s'\n",
			cr.name, eventSpecID, trigger)
	}

	switch trigger {
	case triggerEveryThreshold:
		correlationEvent := cr.createCorrelationEvent(payload.events, false)
		cr.aggregator.feedRoot(correlationEvent, payload.eventSpec, payload.cellKey)
	}
}

// This will be concurrently called by root aggregation cells
func (cr *CorrelationRule) onRootTrigger(trigger string, payload *triggerPayload) {
	if Debug {
		eventSpecID := "-"
		if payload.eventSpec != nil {
			eventSpecID = payload.eventSpec.id
		}
		fmt.Printf(
			"Rule '%s' root-triggered by '%s' on '%s'\n",
			cr.name, eventSpecID, trigger)
	}

	switch trigger {
	case triggerAllThresholdsReached:
		if len(payload.events) > 1 &&
			cr.filter != nil &&
			!cr.filter.Pass(payload.events...) {
			return
		}
	}

	actions, defined := cr.actions[trigger]

	if !defined {
		return
	}

	correlationEvent := cr.createCorrelationEvent(payload.events, true)

	for _, action := range actions {
		action.Take(correlationEvent)
	}

	cr.correlationChain <- correlationEvent
}

func (cr *CorrelationRule) createCorrelationEvent(baseEvents []*normalization.Event, root bool) *normalization.Event {
	event := new(normalization.Event)
	now := time.Now()

	event.ID = GetUUID()
	event.Correlated = true

	baseLen := len(baseEvents)
	event.BaseEventCount = baseLen
	event.CorrelationRuleName = cr.name

	if baseLen == 0 {
		event.StartTime = now
		event.EndTime = now
		return event
	}

	sortEventsByEndTime(baseEvents, false)
	event.BaseEvents = baseEvents

	firstBaseEvent := baseEvents[len(baseEvents)-1]
	lastBaseEvent := baseEvents[0]

	event.StartTime = firstBaseEvent.StartTime
	event.EndTime = lastBaseEvent.EndTime

	for _, field := range cr.aggregation.identicalFields {
		event.SetFieldNoConversion(field, lastBaseEvent.GetFieldNoType(field))
	}

	if root {
		event.BaseEvents = cr.flattenBaseEvents(event.BaseEvents, event.ID)
		event.BaseEventCount = len(event.BaseEvents) - len(cr.eventSpecs)
	} else {
		sumEventFields(cr.aggregation.sumFields, baseEvents, event)
	}

	return event
}

func (cr *CorrelationRule) flattenBaseEvents(baseEvents []*normalization.Event, parentID string) []*normalization.Event {
	allBaseEvents := make([]*normalization.Event, 0)

	for _, baseEvent := range baseEvents {
		baseEvent.ParentID = parentID
		allBaseEvents = append(allBaseEvents, baseEvent)
		if len(baseEvent.BaseEvents) > 0 {
			anotherBaseEvents := cr.flattenBaseEvents(baseEvent.BaseEvents, parentID)
			allBaseEvents = append(allBaseEvents, anotherBaseEvents...)
		}
	}

	return allBaseEvents
}
