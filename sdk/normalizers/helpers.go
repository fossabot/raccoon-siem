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
		event.SetAnyFieldBytes(m.EventField, parsingResult[m.SourceField], m.timeFormat)
	}

	for _, m := range mapping {
		for _, ext := range m.Extra {
			if ext.TriggerField == "" || (ext.TriggerField != "" &&
				bytes.Equal(parsingResult[ext.TriggerField], ext.TriggerValue)) {
				ext.Normalizer.Normalize(parsingResult[m.SourceField], event)
				break
			}
		}
	}

	return event
}
