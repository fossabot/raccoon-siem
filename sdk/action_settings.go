package sdk

type ActionSettings struct {
	SetEventFields *SetEventFieldsActionSettings `yaml:"setEventFields,omitempty"`
	ActiveList     *ActiveListActionSettings     `yaml:"activeList,omitempty"`
}

func (s *ActionSettings) compile() (*actionSpecifications, error) {
	var err error
	spec := new(actionSpecifications)

	// SetEventFields

	if s.SetEventFields != nil {
		spec.setEventFields, err = s.SetEventFields.compile()
		if err != nil {
			return nil, err
		}
	}

	// Active list

	if s.ActiveList != nil {
		spec.activeLists, err = s.ActiveList.compile()
		if err != nil {
			return nil, err
		}
	}

	return spec, nil
}
