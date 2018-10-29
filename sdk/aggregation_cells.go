package sdk

import (
	"time"
)

type aggregationCells struct {
	baseAggregationCells
	data          map[string]*aggregationCell
	root          bool
	rootThreshold int
}

func (acs *aggregationCells) create(key string) *aggregationCell {
	acs.mu.Lock()

	eg := newAggregationCell(acs, key)
	acs.data[key] = eg

	acs.mu.Unlock()
	return eg
}

func (acs *aggregationCells) get(key string) (cell *aggregationCell, ok bool) {
	acs.mu.RLock()
	cell, ok = acs.data[key]
	acs.mu.RUnlock()
	return
}

func (acs *aggregationCells) delete(key string) {
	acs.mu.Lock()
	delete(acs.data, key)
	acs.mu.Unlock()
}

func (acs *aggregationCells) reset() {
	acs.mu.Lock()
	acs.data = make(map[string]*aggregationCell)
	acs.mu.Unlock()
}

func newAggregationCells(
	root bool,
	rootThreshold int,
	eventSpecs []*eventSpecification,
	window time.Duration,
	triggerFunc triggerCallback,
) *aggregationCells {
	acs := new(aggregationCells)
	acs.root = root
	acs.rootThreshold = rootThreshold
	acs.eventSpecs = eventSpecs
	acs.window = window
	acs.triggerFunc = triggerFunc
	acs.reset()
	return acs
}
