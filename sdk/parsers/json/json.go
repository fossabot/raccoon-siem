package json

import (
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/parsers"
)

const (
	dot               = byte('.')
	doubleQuote       = byte('"')
	backSlash         = byte('\\')
	openCurlyBracket  = byte('{')
	closeCurlyBracket = byte('}')
	openBracket       = byte('[')
	closeBracket      = byte(']')
	colon             = byte(':')
	comma             = byte(',')
	space             = byte(' ')
	tab               = byte('\t')
	newLine           = byte('\n')
	carriageReturn    = byte('\r')
	nullByte          = 0
)

const (
	valueKindInvalid = iota
	valueKindObject
	valueKindArray
	valueKindString
	valueKindNumberBoolNull
)

type valueKind uint8

type Config struct {
	parsers.BaseConfig
}

type parser struct {
	cfg Config
}

func (r *parser) ID() string {
	return r.cfg.Name
}

func NewParser(cfg Config) (*parser, error) {
	return &parser{cfg}, nil
}

func (r *parser) Parse(data []byte) (map[string]string, bool) {
	result := make(map[string]string)
	offset := 0
	success := r.getValue(nil, &offset, data, result)
	return result, success
}

func (r *parser) getValue(prefix []byte, dataOffset *int, data []byte, result map[string]string) bool {
	if !r.validJson(data) {
		return false
	}

	for {
		key := nextKey(data, dataOffset)
		if key == nil {
			break
		}

		kind, valueStart := determineValueKind(data, dataOffset)
		if kind == valueKindInvalid {
			break
		}

		if kind == valueKindArray {
			skipValue(kind, data, dataOffset)
			continue
		}

		if kind == valueKindObject {
			r.getValue(append(key, dot), dataOffset, data, result)
			continue
		}

		value := extractValue(kind, data, dataOffset, valueStart)
		if prefix != nil {
			result[string(append(prefix, key...))] = helpers.BytesToString(value)
		} else {
			result[string(key)] = helpers.BytesToString(value)
		}

		if determineIfObjectEnds(data,	 dataOffset) {
			return true
		}
	}

	return true
}

func (r *parser) validJson(data []byte) bool {
	return data[0] == openCurlyBracket && data[len(data) - 1] == closeCurlyBracket
}

func nextKey(data []byte, offset *int) []byte {
	if !goToByte(doubleQuote, data, offset) {
		return nil
	}

	keyStart := *offset + 1

	if !goToByte(doubleQuote, data, offset) {
		return nil
	}

	keyEnd := *offset

	if nextMeaningfulByte(data, offset, false) != colon {
		return nil
	}

	return data[keyStart:keyEnd]
}

func determineIfObjectEnds(data []byte, offset *int) bool {
	switch nextMeaningfulByte(data, offset, false) {
	case closeCurlyBracket:
		*offset -= 1
		return true
	default:
		*offset -= 1
		return false
	}
}

func determineValueKind(data []byte, offset *int) (kind valueKind, pos int) {
	switch nextMeaningfulByte(data, offset, false) {
	case nullByte:
		break
	case openBracket:
		kind = valueKindArray
	case openCurlyBracket:
		kind = valueKindObject
	case doubleQuote:
		kind = valueKindString
	default:
		kind = valueKindNumberBoolNull
		pos = *offset
		return
	}
	pos = *offset + 1
	return
}

func extractValue(kind valueKind, data []byte, offset *int, start int) []byte {
	switch kind {
	case valueKindString:
		skipString(data, offset)
	case valueKindNumberBoolNull:
		skipNumberBoolNull(data, offset)
	default:
		return nil
	}

	if *offset <= len(data) {
		return data[start:*offset]
	}

	return nil
}

func skipValue(kind valueKind, data []byte, offset *int) {
	switch kind {
	case valueKindString:
		skipString(data, offset)
	case valueKindNumberBoolNull:
		skipNumberBoolNull(data, offset)
	default:
		skipComplex(data, offset)
	}
}

func skipString(data []byte, offset *int) {
	goToByte(doubleQuote, data, offset)
}

func skipNumberBoolNull(data []byte, offset *int) {
	for {
		switch nextMeaningfulByte(data, offset, true) {
		case space, tab, newLine, comma, closeCurlyBracket, carriageReturn, nullByte:
			return
		}
	}
}

func skipComplex(data []byte, offset *int) {
	currentLevel := 1
	insideString := false
	for currentLevel > 0 {
		switch nextMeaningfulByte(data, offset, false) {
		case nullByte:
			return
		case doubleQuote:
			insideString = !insideString
		case openCurlyBracket, openBracket:
			if !insideString {
				currentLevel++
			}
		case closeCurlyBracket, closeBracket:
			if !insideString {
				currentLevel--
			}
		}
	}
}

func goToByte(targetByte byte, data []byte, offset *int) bool {
	i := *offset + 1
	for ; i < len(data); i++ {
		currentByte := data[i]
		if currentByte == backSlash {
			i++
		} else if currentByte == targetByte {
			*offset = i
			return true
		}
	}
	*offset = i
	return false
}

func nextMeaningfulByte(data []byte, offset *int, includeSpaces bool) byte {
	i := *offset + 1
	for ; i < len(data); i++ {
		currentByte := data[i]
		switch currentByte {
		case backSlash:
			i++
			continue
		case space, tab, newLine, carriageReturn:
			if includeSpaces {
				*offset = i
				return currentByte
			}
			continue
		default:
			*offset = i
			return currentByte
		}
	}
	*offset = i
	return nullByte
}
