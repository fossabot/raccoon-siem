package sdk

import (
	"time"
)

type simpleAggregationCells struct {
	baseAggregationCells
	data      map[string]*simpleAggregationCell
	sumFields []string
}

func (acs *simpleAggregationCells) create(key string) *simpleAggregationCell {
	acs.mu.Lock()
	defer acs.mu.Unlock()

	eg := newSimpleAggregationCell(acs, key)
	acs.data[key] = eg

	return eg
}

func (acs *simpleAggregationCells) get(key string) (cell *simpleAggregationCell, ok bool) {
	acs.mu.RLock()
	cell, ok = acs.data[key]
	acs.mu.RUnlock()
	return
}

func (acs *simpleAggregationCells) delete(key string) {
	acs.mu.Lock()
	delete(acs.data, key)
	acs.mu.Unlock()
}

func (acs *simpleAggregationCells) reset() {
	acs.mu.Lock()
	acs.data = make(map[string]*simpleAggregationCell)
	acs.mu.Unlock()
}

func newSimpleAggregationCells(
	eventSpecs []*eventSpecification,
	window time.Duration,
	triggerFunc triggerCallback,
	sumFields []string,
) *simpleAggregationCells {
	acs := new(simpleAggregationCells)
	acs.eventSpecs = eventSpecs
	acs.window = window
	acs.triggerFunc = triggerFunc
	acs.sumFields = sumFields
	acs.reset()
	return acs
}
