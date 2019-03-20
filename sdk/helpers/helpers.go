package helpers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strings"
	"time"
)

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}

const (
	sumDelimiter = ','
)

func SumEvents(dst *normalization.Event, src *normalization.Event, fields []string) {
	for _, field := range fields {
		dstValue := dst.GetAnyField(field)
		switch dstValue.(type) {
		case int64:
			srcValue := src.GetIntField(field)
			dst.SetIntField(field, dstValue.(int64)+srcValue)
		case float64:
			newValue := dstValue.(float64) + src.GetFloatField(field)
			dst.SetFloatField(field, newValue)
		case time.Duration:
			newValue := dstValue.(time.Duration) + src.GetDurationField(field)
			dst.SetDurationField(field, newValue)
		case string:
			srcValue, ok := src.GetAnyField(field).(string)
			if ok {
				sb := strings.Builder{}
				if srcValue != "" {
					sb.WriteByte(sumDelimiter)
				}
				sb.WriteString(srcValue)
				dst.SetAnyField(field, sb.String(), normalization.TimeUnitNone)
			}
		default:
			continue
		}
	}
}

func CopyFields(dst *normalization.Event, src *normalization.Event, fields []string) {
	for _, field := range fields {
		srcValue := src.GetAnyField(field)
		switch srcValue.(type) {
		case string:
			dst.SetAnyField(field, srcValue.(string), normalization.TimeUnitNone)
		case int64:
			dst.SetIntField(field, srcValue.(int64))
		case float64:
			dst.SetFloatField(field, srcValue.(float64))
		case time.Duration:
			dst.SetDurationField(field, srcValue.(time.Duration))
		}
	}
}
