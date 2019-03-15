package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"regexp"
	"strings"
)

const (
	smartValueSourceConstant = iota
	smartValueSourceVariable
	smartValueSourceEventField
)

var (
	smartValueTypedExpressionRegexp = regexp.MustCompile(`([^(]+)\(([^)]+)\)`)
)

type smartValue struct {
	value  interface{}
	kind   byte
	source byte
	name   string
}

func (v *smartValue) compile(expr string) *smartValue {
	if v.parseTyped(expr) {
		return v
	}
	v.parseUntyped(expr)
	return v
}

func (v *smartValue) resolve(vars map[string]*variable, event *normalization.Event) (interface{}, error) {
	if v.source == smartValueSourceEventField {
		return event.GetFieldNoType(v.name), nil
	}

	if v.source == smartValueSourceVariable {
		varPtr := vars[v.name]
		val, err := varPtr.getValue(event)

		if err != nil {
			return nil, err
		}

		return normalization.ConvertValue(val, v.kind, normalization.TimeUnitNone), nil
	}

	return v.value, nil
}

func (v *smartValue) parseUntyped(expr string) {
	if yes, kind := v.isEventField(expr); yes {
		v.source = smartValueSourceEventField
		v.kind = kind
		v.name = expr
		return
	}

	if v.isVariable(expr) {
		v.source = smartValueSourceVariable
		v.kind = normalization.FieldTypeString
		v.name = expr[len(variablePrefix):]
		return
	}

	if expr == "false" {
		v.value = false
		v.kind = normalization.FieldTypeBool
	} else if expr == "true" {
		v.value = true
		v.kind = normalization.FieldTypeBool
	} else {
		v.value = expr
		v.kind = normalization.FieldTypeString
	}

	v.source = smartValueSourceConstant
}

func (v *smartValue) parseTyped(expr string) bool {
	matches := smartValueTypedExpressionRegexp.FindStringSubmatch(expr)

	if matches == nil {
		return false
	}

	if !v.setKindFromString(matches[1]) {
		return false
	}

	val := matches[2]

	if v.isVariable(val) {
		v.source = smartValueSourceVariable
		v.name = val[len(variablePrefix):]
	} else {
		v.source = smartValueSourceConstant
		v.value = normalization.ConvertValue(val, v.kind, normalization.TimeUnitNone)
	}

	return true
}

func (v *smartValue) isEventField(s string) (bool, byte) {
	if kind, err := ValidateEventFieldAndGetType(s); err == nil {
		return true, kind
	}
	return false, 0
}

func (v *smartValue) isVariable(s string) bool {
	return strings.HasPrefix(s, variablePrefix)
}

func (v *smartValue) setKindFromString(s string) bool {
	switch s {
	case "string":
		v.kind = normalization.FieldTypeString
	case "int":
		v.kind = normalization.FieldTypeInt
	case "float":
		v.kind = normalization.FieldTypeFloat
	case "bool":
		v.kind = normalization.FieldTypeBool
	case "time":
		v.kind = normalization.FieldTypeTime
	case "duration":
		v.kind = normalization.FieldTypeDuration
	default:
		return false
	}
	return true
}
