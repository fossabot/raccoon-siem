package helpers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strings"
	"time"
)

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
