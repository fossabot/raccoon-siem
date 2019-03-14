package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

type IAction interface {
	Take(event *normalization.Event) error
}

type actionSpecifications struct {
	setEventFields *setEventFieldActionSpecification
	activeLists    *activeListActionSpecification
}

type actionsByTrigger map[string][]IAction

func NewAction(spec *actionSpecifications) IAction {
	switch {
	case spec.setEventFields != nil:
		return newSetEventFieldAction(spec.setEventFields)
	case spec.activeLists != nil:
		return newActiveListAction(spec.activeLists)
	default:
		return nil
	}
}
