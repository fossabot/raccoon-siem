package normalization

import (
	"github.com/araddon/dateparse"
	"strconv"
	"time"
	"unsafe"
)

const (
	TimeUnitSecondsString      = "s"
	TimeUnitMillisecondsString = "ms"
	TimeUnitMicrosecondsString = "us"
	TimeUnitNanosecondsString  = "ns"
)

const (
	TimeUnitNone = iota
	TimeUnitSeconds
	TimeUnitMilliseconds
	TimeUnitMicroseconds
	TimeUnitNanoseconds
)

func BytesToString(input []byte) string {
	return *(*string)(unsafe.Pointer(&input))
}

func BytesToFloat(input string) float64 {
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0
	}
	return float
}

func StringToInt(input string) int64 {
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToBool(input string) bool {
	return input == "true"
}

func StringToDuration(input string, unit byte) time.Duration {
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0
	}
	return DurationFromInt(num, unit)
}

func StringToTime(input string, unit byte) time.Time {
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return time.Unix(0, 0)
	}
	return TimeFromInt(num, unit)
}

func ConvertValue(src interface{}, dstType byte, timeUnit byte) interface{} {
	switch dstType {
	case FieldTypeString:
		return ToString(src)
	case FieldTypeInt:
		return ToInt(src)
	case FieldTypeFloat:
		return ToFloat(src)
	case FieldTypeBool:
		return ToBool(src)
	case FieldTypeTime:
		return ToTime(src, timeUnit)
	case FieldTypeDuration:
		return ToDuration(src, timeUnit)
	default:
		return src
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

func ToInt(src interface{}) int64 {
	switch src.(type) {
	case int64:
		return src.(int64)
	case float64:
		return int64(src.(float64))
	case string:
		out, _ := strconv.ParseInt(src.(string), 10, 64)
		return out
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

func ToFloat(src interface{}) float64 {
	switch src.(type) {
	case float64:
		return src.(float64)
	case int64:
		return float64(src.(int64))
	case string:
		out, _ := strconv.ParseFloat(src.(string), 64)
		return out
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
		out, _ := strconv.ParseBool(src.(string))
		return out
	default:
		return false
	}
}

func ToTime(src interface{}, unit byte) time.Time {
	switch src.(type) {
	case int64:
		return TimeFromInt(src.(int64), unit)
	case float64:
		return TimeFromInt(int64(src.(float64)), unit)
	case string:
		out, _ := dateparse.ParseAny(src.(string))
		return out
	default:
		return time.Time{}
	}
}

func ToDuration(src interface{}, unit byte) time.Duration {
	switch src.(type) {
	case int64:
		return DurationFromInt(src.(int64), unit)
	case float64:
		return DurationFromInt(int64(src.(float64)), unit)
	case string:
		out, _ := time.ParseDuration(src.(string))
		return out
	default:
		return 0
	}
}

func TimeFromInt(v int64, unit byte) time.Time {
	switch unit {
	case TimeUnitSeconds:
		return time.Unix(v, 0)
	case TimeUnitMilliseconds:
		return time.Unix(0, v*1000000)
	case TimeUnitMicroseconds:
		return time.Unix(0, v*1000)
	case TimeUnitNanoseconds:
		return time.Unix(0, v)
	default:
		return time.Unix(0, 0)
	}
}

func DurationFromInt(v int64, unit byte) time.Duration {
	switch unit {
	case TimeUnitSeconds:
		return time.Duration(v) * time.Second
	case TimeUnitMilliseconds:
		return time.Duration(v) * time.Millisecond
	case TimeUnitMicroseconds:
		return time.Duration(v) * time.Microsecond
	default:
		return time.Duration(v)
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
