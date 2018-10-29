package sdk

import "errors"

var (
	errInputTypeUnsupported     = errors.New("unsupported input type")
	errTimeUnitUnsupported      = errors.New("unsupported time unit")
	errCorrelationRulesNotFound = errors.New("no correlation rules found")
	ErrAllParsersFailed         = errors.New("all parsers failed")
	ErrJSONFieldDoesNotExist    = errors.New("json field does not exist")
)

var (
	errDictionaryKeyNotFound = errors.New("dictionary key not found")
	errKeyNotFound           = errors.New("key not found")
	errMalformedInput        = errors.New("malformed input")
)
