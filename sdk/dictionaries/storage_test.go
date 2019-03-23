package dictionaries

import (
	"gotest.tools/assert"
	"testing"
)

const (
	raccoon = "raccoon"
	weird   = "weird"
)

func TestDictionary(t *testing.T) {
	raccoonDict := Config{
		Name: raccoon,
		Data: map[string]interface{}{
			"ERROR": "error",
			"DEBUG": "debug",
			"INFO":  "info",
		},
	}

	weirdDict := Config{
		Name: weird,
		Data: map[string]interface{}{
			"1": "error",
			"2": "debug",
			"3": "info",
		},
	}

	storage := NewStorage([]Config{raccoonDict, weirdDict})

	assert.Equal(t, storage.Get(raccoon, "ERROR"), "error")
	assert.Equal(t, storage.Get(weird, "2"), "debug")
}
