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

func ValidateAndTransformActiveListAction(name string) (byte, error) {
	transformed, ok := activeListActionsTransform[name]

	if !ok {
		return 0, fmt.Errorf("unknown active list action '%s'", name)
	}

	return transformed, nil
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
