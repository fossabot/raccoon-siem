package normalizers

import (
	"gotest.tools/assert"
	"testing"
)

var csvSample = []byte(`CEF:0,енот,threatmanager,1.0,100,detected a \, in message,10,10.0.0.1,blocked,1.1.1.1`)

func TestCSV(t *testing.T) {
	n, err := getCSVNormalizer()
	assert.Equal(t, err, nil)
	event := n.Normalize(csvSample, nil)
	assert.Equal(t, event.UserString1, "CEF:0")
	assert.Equal(t, event.UserString2, "енот")
	assert.Equal(t, event.UserString3, "threatmanager")
	assert.Equal(t, event.UserString4, "1.0")
	assert.Equal(t, event.UserString5, "100")
	assert.Equal(t, event.UserString6, `detected a \, in message`)
	assert.Equal(t, event.UserString7, "10")
	assert.Equal(t, event.UserString8, "10.0.0.1")
	assert.Equal(t, event.Message, "blocked")
	assert.Equal(t, event.DestinationIPAddress, "1.1.1.1")
}

func BenchmarkCSV(b *testing.B) {
	n, _ := getCSVNormalizer()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Normalize(cefSample, nil)
	}
}

func getCSVNormalizer() (INormalizer, error) {
	cfg := Config{
		Name:      "test",
		Kind:      KindCSV,
		Delimiter: ",",
		Mapping: []MappingConfig{
			{SourceField: "0", EventField: "UserString1"},
			{SourceField: "1", EventField: "UserString2"},
			{SourceField: "2", EventField: "UserString3"},
			{SourceField: "3", EventField: "UserString4"},
			{SourceField: "4", EventField: "UserString5"},
			{SourceField: "5", EventField: "UserString6"},
			{SourceField: "6", EventField: "UserString7"},
			{SourceField: "7", EventField: "UserString8"},
			{SourceField: "8", EventField: "Message"},
			{SourceField: "9", EventField: "DestinationIPAddress"},
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return New(cfg)
}
