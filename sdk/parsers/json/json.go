package JSONParser

import (
	"bytes"
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

func GetValue(data []byte, path []byte) []byte {
	pathOffset := 0
	dataOffset := 0
pathLoop:
	for pathOffset < len(path) {
		targetKey := nextPathKey(path, &pathOffset)
		isLastTargetKey := pathOffset >= len(path)
		for {
			key := nextKey(data, &dataOffset)
			if key == nil {
				return nil
			}

			kind, valueStart := determineValueKind(data, &dataOffset)
			if kind == valueKindInvalid {
				return nil
			}

			if !bytes.Equal(targetKey, key) {
				skipValue(kind, data, &dataOffset)
				continue
			}

			if kind == valueKindArray {
				return nil
			}

			if !isLastTargetKey {
				if kind != valueKindObject {
					return nil
				}
				continue pathLoop
			}

			return extractValue(kind, data, &dataOffset, valueStart)
		}
	}
	return nil
}

func nextPathKey(path []byte, offset *int) []byte {
	idx := bytes.IndexByte(path[*offset:], dot)
	if idx == -1 {
		start := *offset
		*offset = len(path)
		return path[start:]
	}
	*offset = idx + 1
	return path[:idx]
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
