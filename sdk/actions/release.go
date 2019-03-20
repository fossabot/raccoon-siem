package actions

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

func Release(cfg ReleaseConfig, correlationEvent *normalization.Event, baseEvents []*normalization.Event) {
	for _, mutateCfg := range cfg.MutateConfigs {
		switch mutateCfg.ValueSourceKind {
		case ValueSourceKindEvent:
			mutateFromEvent(mutateCfg, correlationEvent, baseEvents)
		default:
			correlationEvent.SetAnyField(mutateCfg.Field, mutateCfg.Constant, normalization.TimeUnitNone)
		}
	}
}

func mutateFromEvent(cfg MutateConfig, correlationEvent *normalization.Event, baseEvents []*normalization.Event) {
	var sourceEvent *normalization.Event
	for _, base := range baseEvents {
		if base.Tag == cfg.ValueSourceName {
			sourceEvent = base
			break
		}
	}

	sourceField := cfg.ValueSourceField
	if sourceField == "" {
		sourceField = cfg.Field
	}

	sourceValue := sourceEvent.GetAnyField(sourceField)
	switch sourceValue.(type) {
	case string:
		correlationEvent.SetAnyField(cfg.Field, sourceValue.(string), normalization.TimeUnitNone)
	case int64:
		correlationEvent.SetIntField(cfg.Field, sourceValue.(int64))
	case float64:
		correlationEvent.SetFloatField(cfg.Field, sourceValue.(float64))
	}
}
