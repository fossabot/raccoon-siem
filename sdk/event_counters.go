package sdk

import (
	"sync"
)

func newEventCounters(eventSpecs []*eventSpecification, rootThreshold int) eventCounters {
	ecs := eventCounters{data: make(map[string]*eventCounter)}
	for _, es := range eventSpecs {
		effectiveThreshold := rootThreshold

		if effectiveThreshold == -1 {
			effectiveThreshold = es.threshold
		}

		ecs.data[es.id] = newEventCounter(effectiveThreshold)
	}
	return ecs
}

type eventCounters struct {
	mu   sync.RWMutex
	data map[string]*eventCounter
}

func (ecs *eventCounters) getBySpecID(eventSpecID string) *eventCounter {
	ecs.mu.RLock()
	defer ecs.mu.RUnlock()
	return ecs.data[eventSpecID]
}

func (ecs *eventCounters) checkIfAllThresholdsReached() (reached bool, never bool) {
	ecs.mu.RLock()
	defer ecs.mu.RUnlock()

	for _, counter := range ecs.data {
		if counter.threshold == 0 {
			if counter.eventCount != 0 {
				never = true
				return
			}
		} else if counter.thresholdCount == 0 {
			return
		}
	}

	reached = true
	return
}

func (ecs *eventCounters) reset() {
	ecs.mu.Lock()
	defer ecs.mu.Unlock()

	for _, counter := range ecs.data {
		counter.reset(true)
	}
}
