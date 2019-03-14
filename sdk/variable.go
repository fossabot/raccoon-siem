package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"regexp"
	"strings"
)

const (
	variableSourceActiveListString = "al"
	variableSourceDictionaryString = "dict"
	variableSourceKeySeparator     = "+"
	variablePrefix                 = "$"

	variableSourceActiveList = iota
	variableSourceDictionary
)

var (
	variableExpression   = regexp.MustCompile(`(\S+) = ([^\:]+):([^\:]+):([^\[]+)\[([^\]]*)\]\.?(\S+)?`)
	knownVariableSources = map[string]byte{
		variableSourceActiveListString: variableSourceActiveList,
		variableSourceDictionaryString: variableSourceDictionary,
	}
)

type variable struct {
	name         string
	sourceKind   byte
	sourceAction byte
	sourceName   string
	sourceKeys   []string
	sourceField  string
}

func (v *variable) compile(expr string) (*variable, error) {
	matches := variableExpression.FindStringSubmatch(expr)

	if matches == nil {
		return v, fmt.Errorf("invalid expression '%s", expr)
	}

	// Variable name

	v.name = matches[1]

	if v.name == "" {
		return v, fmt.Errorf("variable must have a name")
	}

	// Variable source kind

	var err error

	v.sourceKind, err = ValidateAndTransformVariableSource(matches[2])

	if err != nil {
		return v, err
	}

	// Variable source name

	v.sourceName = matches[4]

	if v.sourceName == "" {
		return v, fmt.Errorf("variable source must have a name")
	}

	// Variable source action

	action := matches[3]

	switch v.sourceKind {
	case variableSourceActiveList:
		v.sourceAction, err = ValidateAndTransformActiveListAction(action)
		if _, ok := activeListsByName[v.sourceName]; !ok {
			return v, fmt.Errorf("active list '%s' not registered", v.sourceName)
		}
	case variableSourceDictionary:
		v.sourceAction, err = ValidateAndTransformDictionaryAction(action)
		if _, ok := dictionariesByName[v.sourceName]; !ok {
			return v, fmt.Errorf("dictionary '%s' not registered", v.sourceName)
		}
	}

	if err != nil {
		return v, err
	}

	// Variable source keys

	keys := matches[5]

	if keys == "" {
		return nil, fmt.Errorf("variable source keys must be specified")
	}

	parts := strings.Split(keys, variableSourceKeySeparator)
	for _, part := range parts {
		_, err = ValidateEventFieldAndGetType(part)

		if err != nil {
			return v, err
		}

		v.sourceKeys = append(v.sourceKeys, part)
	}

	// Variable source field

	v.sourceField = matches[6]

	if v.sourceKind == variableSourceActiveList &&
		v.sourceAction == actionActiveListKindGet &&
		v.sourceField == "" {
		return v, fmt.Errorf("active list source field must be specified")
	}

	return v, nil
}

func (v *variable) getValue(event *normalization.Event) (interface{}, error) {
	switch v.sourceKind {
	case variableSourceActiveList:
		return v.getValueFromActiveList(event)
	case variableSourceDictionary:
		return v.getValueFromDictionary(event)
	default:
		return "", nil
	}
}

func (v *variable) getValueFromActiveList(event *normalization.Event) (interface{}, error) {
	key := makeActiveListKey(v.sourceName, v.sourceKeys, event)
	al := activeListsByName[v.sourceName]

	switch v.sourceAction {
	default:
		return nil, fmt.Errorf("unsupported active list action")
	case actionActiveListKindCount:
		return al.Count(key)
	case actionActiveListKindSize:
		return al.Size()
	case actionActiveListKindGet:
		resultMap, err := al.Get(key)

		if err != nil {
			return nil, err
		}

		val, ok := resultMap[v.sourceField]

		if !ok {
			return nil, fmt.Errorf("active list record field '%s' does not exist", v.sourceField)
		}

		return string(val.([]byte)), nil
	}
}

func (v *variable) getValueFromDictionary(event *normalization.Event) (interface{}, error) {
	dict := dictionariesByName[v.sourceName]
	key := event.GetFieldNoType(v.sourceKeys[0])

	val, ok := dict.data[key]

	if !ok {
		return nil, errDictionaryKeyNotFound
	}

	return val, nil
}
