package sdk

import (
	"fmt"
	"regexp"
)

const (
	actionActiveListKindAddString    = "add"
	actionActiveListKindGetString    = "get"
	actionActiveListKindDeleteString = "delete"
	actionActiveListKindCountString  = "count"
	actionActiveListKindSizeString   = "size"
)

var activeListActionsTransform = map[string]byte{
	actionActiveListKindAddString:    actionActiveListKindAdd,
	actionActiveListKindGetString:    actionActiveListKindGet,
	actionActiveListKindDeleteString: actionActiveListKindDelete,
	actionActiveListKindCountString:  actionActiveListKindCount,
	actionActiveListKindSizeString:   actionActiveListKindSize,
}

type ActiveListActionSettings struct {
	Name      string   `yaml:"name,omitempty"`
	Kind      string   `yaml:"kind,omitempty"`
	KeyFields []string `yaml:"keyFields,omitempty"`
	Fields    []string `yaml:"fields,omitempty"`
}

func (s *ActiveListActionSettings) compile() (*activeListActionSpecification, error) {
	spec := new(activeListActionSpecification)

	// Kind

	transformedKind, err := ValidateAndTransformActiveListAction(s.Kind)

	if err != nil {
		return nil, err
	}

	spec.kind = transformedKind

	// Name

	if s.Name == "" {
		return nil, fmt.Errorf("target active list name is required")
	}

	if _, ok := activeListsByName[s.Name]; !ok {
		return nil, fmt.Errorf("unknown active list '%s'", s.Name)
	}

	spec.name = s.Name

	// Key fields

	if len(s.KeyFields) == 0 {
		return nil, fmt.Errorf("at least 1 key field must be specified")
	}

	for _, kf := range s.KeyFields {
		if _, err := ValidateEventFieldAndGetType(kf); err != nil {
			return nil, err
		}
	}

	spec.keyFields = s.KeyFields

	// Field expressions

	if len(s.Fields) == 0 {
		return nil, fmt.Errorf("at least 1 active list field must be specified")
	}

	re := regexp.MustCompile(`(\S+) = (\S+)`)

	for _, exp := range s.Fields {
		compiledField, err := s.compileField(exp, re)

		if err != nil {
			return nil, err
		}

		spec.fields = append(spec.fields, compiledField)
	}

	return spec, nil
}

func (s *ActiveListActionSettings) compileField(
	exp string,
	re *regexp.Regexp,
) (*activeListActionFieldSpecification, error) {
	f := new(activeListActionFieldSpecification)
	matches := re.FindStringSubmatch(exp)

	if matches == nil {
		return nil, fmt.Errorf("invalid active list field expression '%s", exp)
	}

	f.eventField = matches[2]
	if _, err := ValidateEventFieldAndGetType(f.eventField); err != nil {
		return nil, err
	}

	f.listField = matches[1]

	return f, nil
}
