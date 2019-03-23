package regexp

import (
	"regexp"
)

func Parse(data []byte, expressions []*regexp.Regexp) (map[string][]byte, bool) {
	for _, e := range expressions {
		if match := e.FindSubmatch(data); match != nil {
			output := make(map[string][]byte)
			for i, field := range e.SubexpNames() {
				if i > 0 {
					output[field] = match[i]
				}
			}
			return output, true
		}
	}
	return nil, false
}
