package activeLists

import "strings"

func MakeFinalKey(list, key string) string {
	sb := strings.Builder{}
	sb.WriteString(AlNamePrefix)
	sb.WriteString(list)
	sb.WriteByte(':')
	sb.WriteString(key)
	return sb.String()
}
