package sdk

import (
	"fmt"
	"time"
)

type ICorrelationRule interface {
	ID() string
	Run()
	Feed(event *Event)
}

type CorrelationRule struct {
	BaseRule
	aggregator       *aggregator
	filter           IFilter
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

func (cr *CorrelationRule) Feed(event *Event) {
	for _, eventSpec := range cr.eventSpecs {
		if eventSpec.filter.Pass([]*Event{event}) {
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
			!cr.filter.Pass(payload.events) {
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

func (cr *CorrelationRule) createCorrelationEvent(baseEvents []*Event, root bool) *Event {
	event := new(Event)
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
	event.baseEvents = baseEvents

	firstBaseEvent := baseEvents[len(baseEvents)-1]
	lastBaseEvent := baseEvents[0]

	event.StartTime = firstBaseEvent.StartTime
	event.EndTime = lastBaseEvent.EndTime

	for _, field := range cr.aggregation.identicalFields {
		event.SetFieldNoConversion(field, lastBaseEvent.GetFieldNoType(field))
	}

	if root {
		event.baseEvents = cr.flattenBaseEvents(event.baseEvents, event.ID)
		event.BaseEventCount = len(event.baseEvents) - len(cr.eventSpecs)
	} else {
		sumEventFields(cr.aggregation.sumFields, baseEvents, event)
	}

	return event
}

func (cr *CorrelationRule) flattenBaseEvents(baseEvents []*Event, parentID string) []*Event {
	allBaseEvents := make([]*Event, 0)

	for _, baseEvent := range baseEvents {
		baseEvent.ParentID = parentID
		allBaseEvents = append(allBaseEvents, baseEvent)
		if len(baseEvent.baseEvents) > 0 {
			anotherBaseEvents := cr.flattenBaseEvents(baseEvent.baseEvents, parentID)
			allBaseEvents = append(allBaseEvents, anotherBaseEvents...)
		}
	}

	return allBaseEvents
}
