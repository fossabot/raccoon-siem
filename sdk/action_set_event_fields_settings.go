package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"regexp"
)

type SetEventFieldsActionSettings []string

func (s SetEventFieldsActionSettings) compile() (*setEventFieldActionSpecification, error) {
	spec := new(setEventFieldActionSpecification)

	if len(s) == 0 {
		return nil, fmt.Errorf("at least 1 field must be specified")
	}

	re := regexp.MustCompile(`(\S+) = (.+)`)

	for _, fieldExp := range s {
		af, err := s.compileField(fieldExp, re)

		if err != nil {
			return nil, err
		}

		spec.fields = append(spec.fields, af)
	}

	return spec, nil
}

func (s SetEventFieldsActionSettings) compileField(
	expression string,
	re *regexp.Regexp,
) (*setEventFieldsActionField, error) {
	af := new(setEventFieldsActionField)
	matches := re.FindStringSubmatch(expression)

	if matches == nil {
		return nil, fmt.Errorf("invalid field/value expression '%s'", expression)
	}

	af.name = matches[1]
	eventFieldType, err := ValidateEventFieldAndGetType(af.name)

	if err != nil {
		return nil, err
	}

	af.value = normalization.ConvertValue(matches[2], eventFieldType, normalization.TimeUnitNone)

	return af, nil
}
