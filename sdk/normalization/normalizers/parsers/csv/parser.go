package csv

import (
	"bytes"
	"encoding/csv"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers"
	"strconv"
)

func Parse(data []byte, delimiter byte, callback parsers.Callback) bool {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.Comma = rune(delimiter)
	reader.TrimLeadingSpace = true
	reader.LazyQuotes = true
	reader.ReuseRecord = true

	rec, err := reader.Read()
	if err != nil {
		return false
	}

	for i := range rec {
		callback(strconv.Itoa(i), []byte(rec[i]))
	}

	return true
}
