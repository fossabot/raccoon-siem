package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

const (
	actionActiveListKindAdd = iota
	actionActiveListKindGet
	actionActiveListKindDelete
	actionActiveListKindCount
	actionActiveListKindSize
)

func newActiveListAction(spec *activeListActionSpecification) IAction {
	return &activeListAction{spec: spec}
}

type activeListAction struct {
	spec *activeListActionSpecification
}

type activeListActionSpecification struct {
	kind      byte
	name      string
	keyFields []string
	fields    []*activeListActionFieldSpecification
}

type activeListActionFieldSpecification struct {
	listField  string
	eventField string
}

func (a *activeListAction) Take(event *normalization.Event) error {
	targetAL := activeListsByName[a.spec.name]

	var err error
	var retValue interface{}

	recordKey := makeActiveListKey(targetAL.spec.name, a.spec.keyFields, event)

	switch a.spec.kind {
	case actionActiveListKindAdd:
		err = targetAL.Add(recordKey, a.makeALValuesMap(event))
	case actionActiveListKindGet:
		retValue, err = targetAL.Get(recordKey)
		if err == nil {
			a.setEventFieldsFromALMap(event, retValue.(alMultiValueType))
		}
	case actionActiveListKindDelete:
		err = targetAL.Delete(recordKey)
	case actionActiveListKindCount:
		retValue, err = targetAL.Count(recordKey)
		if err == nil {
			event.SetField(a.spec.fields[0].eventField, retValue.(int64), normalization.TimeUnitNone)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (a *activeListAction) makeALValuesMap(event *normalization.Event) alMultiValueType {
	m := make(alMultiValueType)

	for _, f := range a.spec.fields {
		m[f.listField] = event.GetFieldNoType(f.eventField)
	}

	return m
}

func (a *activeListAction) setEventFieldsFromALMap(event *normalization.Event, m alMultiValueType) {
	for _, f := range a.spec.fields {
		val, ok := m[f.listField]
		if ok {
			event.SetField(f.eventField, val, normalization.TimeUnitNone)
		}
	}
}
