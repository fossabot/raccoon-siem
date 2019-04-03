package normalization

import (
	"github.com/araddon/dateparse"
	"strconv"
	"strings"
	"unsafe"
)

func BytesToString(input []byte) string {
	return *(*string)(unsafe.Pointer(&input))
}

func StringToInt(input string) int64 {
	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToFloat(input string) float64 {
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToTime(input string) int64 {
	t, err := dateparse.ParseAny(strings.TrimSpace(input))
	if err != nil {
		return 0
	}
	return t.UnixNano() / 1000
}

func StringToBool(input string) bool {
	return strings.TrimSpace(input) == "true"
}

func ToFieldType(field string, v interface{}) interface{} {
	fv := new(Event).GetAnyField(field)
	switch fv.(type) {
	case string:
		return ToString(v)
	case int64:
		return ToInt64(v)
	case float64:
		return ToFloat64(v)
	case bool:
		return ToBool(v)
	default:
		return nil
	}
}

func ToString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int64:
		return strconv.FormatInt(src.(int64), 10)
	case float64:
		return strconv.FormatFloat(src.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(src.(bool))
	default:
		return ""
	}
}

func ToInt64(src interface{}) int64 {
	switch src.(type) {
	case int64:
		return src.(int64)
	case float64:
		return int64(src.(float64))
	case string:
		return StringToInt(src.(string))
	case bool:
		if src.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func ToFloat64(src interface{}) float64 {
	switch src.(type) {
	case float64:
		return src.(float64)
	case int64:
		return float64(src.(int64))
	case string:
		return StringToFloat(src.(string))
	case bool:
		if src.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func ToBool(src interface{}) bool {
	switch src.(type) {
	case bool:
		return src.(bool)
	case int64:
		return src.(int64) > 0
	case float64:
		return src.(float64) > 0
	case string:
		return StringToBool(src.(string))
	default:
		return false
	}
}

func To64Bits(i interface{}) interface{} {
	switch i.(type) {
	case int:
		return int64(i.(int))
	case int8:
		return int64(i.(int8))
	case int16:
		return int64(i.(int16))
	case int32:
		return int64(i.(int32))
	case uint:
		return uint64(i.(uint))
	case uint8:
		return uint64(i.(uint8))
	case uint16:
		return uint64(i.(uint16))
	case uint32:
		return uint64(i.(uint32))
	case float32:
		return float64(i.(float32))
	default:
		return i
	}
}
