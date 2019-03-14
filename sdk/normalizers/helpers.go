package normalizers

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

func normalize(
	parsingResult map[string][]byte,
	mapping []MappingConfig,
	event *normalization.Event,
) *normalization.Event {
	if event == nil {
		event = new(normalization.Event)
	}

	for _, m := range mapping {
		if m.Extra == nil {
			event.Set(m.EventField, parsingResult[m.SourceField], m.timeFormat)
		}
	}

	for _, m := range mapping {
		if m.Extra != nil {
			sourceValue := parsingResult[m.SourceField]
			event.Set(m.EventField, sourceValue, m.timeFormat)
			m.Extra.Normalizer.Normalize(sourceValue, event)
		}
	}

	return event
}
