package normalizers

import (
	"bytes"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

func normalize(
	parsingResult map[string][]byte,
	mapping []MappingConfig,
	event *normalization.Event,
) *normalization.Event {
	if event == nil {
		event = new(normalization.Event)
	}

	for _, m := range mapping {
		sourceValue := parsingResult[m.SourceField]
		if sourceValue != nil {
			event.SetAnyFieldBytes(m.EventField, sourceValue)
			event.FieldsNormalized++
			if m.labelField != "" {
				event.SetAnyField(m.labelField, m.Label)
			}
		}
	}

	for _, m := range mapping {
		for _, ext := range m.Extra {
			if ext.TriggerField == "" || (ext.TriggerField != "" &&
				bytes.Equal(parsingResult[ext.TriggerField], ext.triggerValue)) {
				ext.normalizer.Normalize(parsingResult[m.SourceField], event)
				break
			}
		}
	}

	return event
}
