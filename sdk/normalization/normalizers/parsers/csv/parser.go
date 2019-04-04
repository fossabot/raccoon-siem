package csv

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers"
	"strconv"
)

const (
	bs = '\\'
)

func Parse(data []byte, delimiter byte, callback parsers.Callback) bool {
	delimitersFound := 0
	pos := 0
	valueStart := pos

	for ; pos < len(data); pos++ {
		if data[pos] == delimiter && data[pos-1] != bs {
			callback(strconv.Itoa(delimitersFound), data[valueStart:pos])
			delimitersFound++
			valueStart = pos + 1
		}
	}

	if delimitersFound > 0 {
		callback(strconv.Itoa(delimitersFound), data[valueStart:pos])
		return true
	}

	return false
}
