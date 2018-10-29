package sdk

import "sync"

func newEventUniqueHashes() eventUniqueHashes {
	return eventUniqueHashes{data: make(map[string]map[string]bool)}
}

type eventUniqueHashes struct {
	mu   sync.RWMutex
	data map[string]map[string]bool
}

func (h *eventUniqueHashes) addIfNotExists(specID string, key string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	hashes, ok := h.data[specID]

	if !ok {
		hashes = make(map[string]bool)
		hashes[key] = true
		h.data[specID] = hashes
		return true
	}

	if _, ok := hashes[key]; ok {
		return false
	}

	hashes[key] = true
	h.data[specID] = hashes

	return true
}

func (h *eventUniqueHashes) reset(specID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if specID != anySpecID {
		h.data[specID] = make(map[string]bool)
		return
	}

	h.data = make(map[string]map[string]bool)
}
