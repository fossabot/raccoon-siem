package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strings"
)

func makeKey(keyFields []string, event *normalization.Event) string {
	key := strings.Builder{}
	for _, field := range keyFields {
		key.WriteString(normalization.ToString(event.GetAnyField(field)))
	}
	return key.String()
}
