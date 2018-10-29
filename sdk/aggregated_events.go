package sdk

import (
	"sync"
)

const maxCapacity = 1024

func newAggregatedEvents() aggregatedEvents {
	return aggregatedEvents{
		data: make(map[string][]*Event),
	}
}

type aggregatedEvents struct {
	mu   sync.RWMutex
	data map[string][]*Event
}

func (ae *aggregatedEvents) add(key string, e *Event) {
	ae.mu.Lock()
	defer ae.mu.Unlock()

	slice, ok := ae.data[key]

	if !ok {
		slice = make([]*Event, 0)
	}

	if len(slice) == maxCapacity {
		slice = append(slice[1:], e)
	} else {
		slice = append(slice, e)
	}

	ae.data[key] = slice
}

func (ae *aggregatedEvents) sum(key string, fields []string, e *Event, threshold int) bool {
	ae.mu.Lock()
	defer ae.mu.Unlock()

	slice, ok := ae.data[key]

	if !ok {
		e.AggregatedEventCount = 1
		slice = []*Event{e}
		ae.data[key] = slice
		return threshold == 1
	}

	targetEvent := slice[0]
	targetEvent.AggregatedEventCount++
	sumEventFields(fields, []*Event{e}, targetEvent)

	if targetEvent.AggregatedEventCount == threshold {
		targetEvent.EndTime = e.EndTime
		return true
	}

	return false
}

func (ae *aggregatedEvents) get(key string) []*Event {
	ae.mu.RLock()
	defer ae.mu.RUnlock()

	slice, ok := ae.data[key]

	if !ok {
		return make([]*Event, 0)
	}

	return slice
}

func (ae *aggregatedEvents) getAll() (result []*Event) {
	ae.mu.RLock()
	defer ae.mu.RUnlock()

	for _, v := range ae.data {
		result = append(result, v...)
	}

	return
}

func (ae *aggregatedEvents) reset(specID string) {
	ae.mu.Lock()
	defer ae.mu.Unlock()

	if specID != anySpecID {
		delete(ae.data, specID)
		return
	}

	ae.data = make(map[string][]*Event)
}
