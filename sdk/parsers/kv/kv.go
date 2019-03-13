package kv

import (
	"errors"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

const (
	space = byte(' ')
	bs    = '\\'
)

type Parser struct {
	pairSeparator byte
	kvSeparator byte
}

func NewParser(pairSeparator byte, kvSeparator byte) (*Parser, error) {
	if pairSeparator == kvSeparator {
		return nil, errors.New("kv and pair separators must be different")
	}

	parser := Parser{
		pairSeparator: pairSeparator,
		kvSeparator:   kvSeparator,
	}

	return &parser, nil
}

func (rec *Parser) Parse(data []byte) (map[string]string, bool) {
	result := make(map[string]string)

	var key []byte
	var start = 0
	var end = 0

	onValue := false
	lookForValue := false
	for i := range data {

		// Separator between key and value was met
		if data[i] == rec.kvSeparator && !lookForValue && data[i-1] != bs {
			// Save current key
			key = data[start:end]
			// Wait for value now
			lookForValue = true
			// Value will start with next char
			start = i + 1
			// Mark that we have passed key
			onValue = false

			continue
		}

		// Separator between pairs of "key-value" was met
		if data[i] == rec.pairSeparator && lookForValue && data[i-1] != bs {
			// Save current value to map with early saved key
			result[helpers.BytesToString(key)] = helpers.BytesToString(data[start:end])
			// Wait for next key now
			lookForValue = false
			// Key will start with next char
			start = i + 1
			// Mark that we have passed value
			onValue = false

			continue
		}

		// We have met not space character
		if data[i] != space {
			// Shift end index and mark that we have reached value or key
			end = i + 1
			onValue = true
		}

		// Shift start index if we are not on value slice
		if !onValue {
			start = i + 1
		}

	}

	// Input ends. If we still waiting for value, save value
	if lookForValue {
		result[helpers.BytesToString(key)] = helpers.BytesToString(data[start:end])
	}

	return result, true
}

