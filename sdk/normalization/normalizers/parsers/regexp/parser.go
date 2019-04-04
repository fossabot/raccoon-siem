package regexp

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers"
	"regexp"
)

func Parse(data []byte, expressions []*regexp.Regexp, callback parsers.Callback) bool {
	for _, e := range expressions {
		if match := e.FindSubmatch(data); match != nil {
			for i, field := range e.SubexpNames() {
				if i > 0 {
					callback(field, match[i])
				}
			}
			return true
		}
	}
	return false
}
