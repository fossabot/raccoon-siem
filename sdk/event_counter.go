package sdk

import "sync"

func newEventCounter(threshold int) *eventCounter {
	return &eventCounter{threshold: threshold}
}

type eventCounter struct {
	mu             sync.RWMutex
	eventCount     int
	thresholdCount int
	threshold      int
}

func (c *eventCounter) increment() (int, int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var reached bool
	c.eventCount += 1

	if c.threshold != 0 {
		reached = c.eventCount%c.threshold == 0
	}

	if reached {
		c.thresholdCount += 1
	}

	return c.eventCount, c.thresholdCount, reached
}

func (c *eventCounter) reset(withThresholdCount bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.eventCount = 0
	if withThresholdCount {
		c.thresholdCount = 0
	}
}
