package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

func ValidateEventFieldAndGetType(name string) (byte, error) {
	fType, ok := normalization.EventFieldTypeByName[name]

	if !ok {
		return 0, fmt.Errorf("unknown or protected event field '%s'", name)
	}

	return fType, nil
}

func ValidateTrigger(kind string) error {
	_, ok := knownTriggers[kind]

	if !ok {
		return fmt.Errorf("unknown trigger '%s'", kind)
	}

	return nil
}

func ValidateDestination(name string) error {
	if _, ok := knownDestinations[name]; !ok {
		return fmt.Errorf("unknown destination '%s'", name)
	}
	return nil
}

func ValidateAndTransformActiveListAction(name string) (byte, error) {
	transformed, ok := activeListActionsTransform[name]

	if !ok {
		return 0, fmt.Errorf("unknown active list action '%s'", name)
	}

	return transformed, nil
}

func ValidateAndTransformFilterOperator(v string) (byte, error) {
	switch v {
	case opEQString:
		return opEQ, nil
	case opNEQString:
		return opNEQ, nil
	case opGTString:
		return opGT, nil
	case opLTString:
		return opLT, nil
	case opLTorEQString:
		return opLTorEQ, nil
	case opGTorEQString:
		return opGTorEQ, nil
	default:
		return 0, fmt.Errorf("unknown children operator '%s", v)
	}
}

func ValidateAndTransformTimeUnit(v string) (byte, error) {
	switch v {
	case normalization.TimeUnitSecondsString:
		return normalization.TimeUnitSeconds, nil
	case normalization.TimeUnitMillisecondsString:
		return normalization.TimeUnitMilliseconds, nil
	case normalization.TimeUnitMicrosecondsString:
		return normalization.TimeUnitMicroseconds, nil
	case normalization.TimeUnitNanosecondsString:
		return normalization.TimeUnitNanoseconds, nil
	default:
		return normalization.TimeUnitNone, nil
	}
}

func ValidateAndTransformVariableSource(s string) (byte, error) {
	b, ok := knownVariableSources[s]

	if !ok {
		return 0, fmt.Errorf("unknown variable source '%s'", s)
	}

	return b, nil
}

func ValidateAndTransformDictionaryAction(name string) (byte, error) {
	transformed, ok := knownDictionaryActions[name]

	if !ok {
		return 0, fmt.Errorf("unknown dictionary action '%s'", name)
	}

	return transformed, nil
}
