package sdk

import (
	"time"
)

type aggregationCell struct {
	baseAggregationCell
	container    *aggregationCells
	counters     eventCounters
	uniqueHashes eventUniqueHashes
}

func (ac *aggregationCell) Put(event *Event, uniqueKey string, eventSpec *eventSpecification) {
	ac.mu.Lock()

	resetSpec := ""
	triggers := make([]string, 0, 8)
	defer ac.fireTriggersResetsAndUnlock(&triggers, &resetSpec, eventSpec, event)

	if uniqueKey != defaultEventFieldsHash {
		if !ac.uniqueHashes.addIfNotExists(eventSpec.id, uniqueKey) {
			return
		}
	}

	event.CorrelatorEventSpecID = eventSpec.id

	ac.aggregatedEvents.add(eventSpec.id, event)
	eventCount, thresholdCount, thresholdReached := ac.counters.getBySpecID(eventSpec.id).increment()

	// Every partial event trigger
	triggers = append(triggers, triggerEveryEvent)

	if eventCount == 1 {
		// First partial event trigger
		triggers = append(triggers, triggerFirstEvent)
	} else {
		// Subsequent partial event trigger
		triggers = append(triggers, triggerSubsequentEvents)
	}

	if thresholdCount == 1 {
		// First partial threshold trigger
		triggers = append(triggers, triggerFirstThreshold)
	} else if thresholdCount > 1 {
		// Subsequent partial threshold trigger
		triggers = append(triggers, triggerSubsequentThresholds)
	}

	// If any partial threshold was reached (first or subsequent)
	if thresholdReached {
		// Every partial threshold trigger
		triggers = append(triggers, triggerEveryThreshold)
		if !ac.container.root {
			resetSpec = eventSpec.id
		}
	}

	if ac.container.root {
		reached, neverWill := ac.counters.checkIfAllThresholdsReached()
		if reached {
			// All thresholds reached trigger
			triggers = append(triggers, triggerAllThresholdsReached)
			resetSpec = anySpecID
		} else if neverWill {
			resetSpec = anySpecID
		}
	}
}

func (ac *aggregationCell) fireTriggersResetsAndUnlock(
	triggers *[]string,
	resetSpec *string,
	eventSpec *eventSpecification,
	event *Event,
) {
	for _, t := range *triggers {
		ac.callTrigger(t, eventSpec, event)
	}

	if *resetSpec != "" {
		ac.reset(*resetSpec)
	}

	ac.mu.Unlock()
}

func (ac *aggregationCell) callTrigger(trigger string, eventSpec *eventSpecification, event *Event) {
	payload := &triggerPayload{
		eventSpec:  eventSpec,
		eventSpecs: ac.container.eventSpecs,
		cellKey:    ac.key,
	}

	if eventSpec != nil {
		payload.counter = ac.counters.getBySpecID(eventSpec.id)
	}

	switch trigger {
	case triggerEveryEvent, triggerFirstEvent, triggerSubsequentEvents:
		payload.events = []*Event{event}
	case triggerTimeout:
		payload.events = make([]*Event, 0)
	case triggerAllThresholdsReached:
		payload.events = ac.aggregatedEvents.getAll()
	default:
		payload.events = ac.aggregatedEvents.get(eventSpec.id)
	}

	ac.container.triggerFunc(trigger, payload)
}

func (ac *aggregationCell) reset(specID string) {
	ac.aggregatedEvents.reset(specID)
	ac.uniqueHashes.reset(specID)
	if specID == anySpecID {
		ac.stopTimer()
		ac.counters.reset()
		ac.container.delete(ac.key)
	}
}

func newAggregationCell(container *aggregationCells, key string) *aggregationCell {
	ac := new(aggregationCell)
	ac.key = key
	ac.container = container
	ac.aggregatedEvents = newAggregatedEvents()
	ac.counters = newEventCounters(container.eventSpecs, container.rootThreshold)
	ac.uniqueHashes = newEventUniqueHashes()

	if container.window > 0 {
		ac.timer = time.AfterFunc(container.window, func() {
			ac.callTrigger(triggerTimeout, nil, nil)
			ac.reset(anySpecID)
		})
	}

	return ac
}
