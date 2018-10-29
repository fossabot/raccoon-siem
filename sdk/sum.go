package sdk

import (
	"time"
)

func sumEventFields(fields []string, events []*Event, targetEvent *Event) {
	for _, f := range fields {
		fieldValue := targetEvent.GetFieldNoType(f)
		switch fieldValue.(type) {
		case int64:
			targetEvent.SetFieldNoConversion(f, sumIntFields(events, f))
		case float64:
			targetEvent.SetFieldNoConversion(f, sumFloatFields(events, f))
		case time.Duration:
			targetEvent.SetFieldNoConversion(f, sumDurationFields(events, f))
		default:
			continue
		}
	}
}

func sumIntFields(events []*Event, field string) (sum int64) {
	for _, e := range events {
		v := e.GetFieldNoType(field).(int64)
		sum += v
	}
	return
}

func sumFloatFields(events []*Event, field string) (sum float64) {
	for _, e := range events {
		v := e.GetFieldNoType(field).(float64)
		sum += v
	}
	return
}

func sumDurationFields(events []*Event, field string) (sum time.Duration) {
	for _, e := range events {
		v := e.GetFieldNoType(field).(time.Duration)
		sum += v
	}
	return
}
