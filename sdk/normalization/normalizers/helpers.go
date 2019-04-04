package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers"
)

func extraNormalize(event *normalization.Event, extra []ExtraConfig) *normalization.Event {
	for i := range extra {
		ext := &extra[i]
		if event.GetAnyField(ext.ConditionEventField) == ext.ConditionValue {
			ext.normalizer.Normalize(
				[]byte(normalization.ToString(event.GetAnyField(ext.SourceEventField))),
				event,
			)
		}
	}
	return event
}

func parserCallbackGenerator(mapping map[string]MappingConfig, event *normalization.Event) parsers.Callback {
	return func(key string, value []byte) {
		m, ok := mapping[key]
		if ok {
			event.SetAnyFieldBytes(m.EventField, value)
			event.FieldsNormalized++
			if m.labelField != "" {
				event.SetAnyField(m.labelField, m.Label)
			}
		}
	}
}

func eventOrNil(event *normalization.Event, eventJustCreated bool) *normalization.Event {
	if eventJustCreated {
		return nil
	}
	return event
}

func createEventIfNil(event *normalization.Event) (*normalization.Event, bool) {
	if event == nil {
		return new(normalization.Event), true
	}
	return event, false
}

func groupMappingBySourceField(mapping []MappingConfig) map[string]MappingConfig {
	grouped := make(map[string]MappingConfig)
	for _, m := range mapping {
		grouped[m.SourceField] = m
	}
	return grouped
}
