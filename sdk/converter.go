package sdk

import (
	"github.com/araddon/dateparse"
	"strconv"
	"time"
)

const (
	timeUnitSecondsString      = "s"
	timeUnitMillisecondsString = "ms"
	timeUnitMicrosecondsString = "us"
	timeUnitNanosecondsString  = "ns"
)

const (
	timeUnitNone = iota
	timeUnitSeconds
	timeUnitMilliseconds
	timeUnitMicroseconds
	timeUnitNanoseconds
)

func convertValue(src interface{}, dstType byte, timeUnit byte) interface{} {
	switch dstType {
	case fieldTypeString:
		return toString(src)
	case fieldTypeInt:
		return toInt(src)
	case fieldTypeFloat:
		return toFloat(src)
	case fieldTypeBool:
		return toBool(src)
	case fieldTypeTime:
		return toTime(src, timeUnit)
	case fieldTypeDuration:
		return toDuration(src, timeUnit)
	default:
		return src
	}
}

func toString(src interface{}) string {
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

func toInt(src interface{}) int64 {
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

func toFloat(src interface{}) float64 {
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

func toBool(src interface{}) bool {
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

func toTime(src interface{}, unit byte) time.Time {
	switch src.(type) {
	case int64:
		return timeFromInt(src.(int64), unit)
	case float64:
		return timeFromInt(int64(src.(float64)), unit)
	case string:
		out, _ := dateparse.ParseAny(src.(string))
		return out
	default:
		return time.Time{}
	}
}

func toDuration(src interface{}, unit byte) time.Duration {
	switch src.(type) {
	case int64:
		return durationFromInt(src.(int64), unit)
	case float64:
		return durationFromInt(int64(src.(float64)), unit)
	case string:
		out, _ := time.ParseDuration(src.(string))
		return out
	default:
		return 0
	}
}

func timeFromInt(v int64, unit byte) time.Time {
	switch unit {
	case timeUnitSeconds:
		return time.Unix(v, 0)
	case timeUnitMilliseconds:
		return time.Unix(0, v*1000000)
	case timeUnitMicroseconds:
		return time.Unix(0, v*1000)
	case timeUnitNanoseconds:
		return time.Unix(0, v)
	default:
		return time.Unix(0, 0)
	}
}

func durationFromInt(v int64, unit byte) time.Duration {
	switch unit {
	case timeUnitSeconds:
		return time.Duration(v) * time.Second
	case timeUnitMilliseconds:
		return time.Duration(v) * time.Millisecond
	case timeUnitMicroseconds:
		return time.Duration(v) * time.Microsecond
	default:
		return time.Duration(v)
	}
}

func to64Bits(i interface{}) interface{} {
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
