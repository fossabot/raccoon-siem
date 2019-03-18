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
		if m.Extra != nil {
			if m.Extra.TriggerField == "" || (m.Extra.TriggerField != "" &&
				bytes.Equal(parsingResult[m.Extra.TriggerField], m.Extra.TriggerValue)) {
				m.Extra.Normalizer.Normalize(parsingResult[m.SourceField], event)
			}
		}
	}

	return event
}
