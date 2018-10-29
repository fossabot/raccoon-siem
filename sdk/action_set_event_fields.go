package sdk

func newSetEventFieldAction(spec *setEventFieldActionSpecification) IAction {
	return &setEventFieldsAction{spec: spec}
}

type setEventFieldActionSpecification struct {
	fields []*setEventFieldsActionField
}

type setEventFieldsActionField struct {
	name  string
	value interface{}
}

type setEventFieldsAction struct {
	spec *setEventFieldActionSpecification
}

func (a *setEventFieldsAction) Take(event *Event) error {
	for _, field := range a.spec.fields {
		event.SetFieldNoConversion(field.name, field.value)
	}
	return nil
}
