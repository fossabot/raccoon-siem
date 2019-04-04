package normalizers

import (
	"gotest.tools/assert"
	"testing"
)

var cefSample = []byte(`CEF:0|енот|threatmanager|1.0|100|detected a \| in message|10|src=10.0.0.1 act=blocked message dst=1.1.1.1`)

func TestCEF(t *testing.T) {
	n, err := getCEFNormalizer()
	assert.Equal(t, err, nil)
	event := n.Normalize(cefSample, nil)
	assert.Equal(t, event.UserString1, "енот")
	assert.Equal(t, event.UserString2, "threatmanager")
	assert.Equal(t, event.UserString3, "1.0")
	assert.Equal(t, event.UserString4, "100")
	assert.Equal(t, event.UserString5, `detected a \| in message`)
	assert.Equal(t, event.UserString6, "10")
	assert.Equal(t, event.UserString7, "10.0.0.1")
	assert.Equal(t, event.UserString8, "blocked message")
	assert.Equal(t, event.DestinationIPAddress, "1.1.1.1")
}

func BenchmarkCEF(b *testing.B) {
	n, _ := getKVNormalizer()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Normalize(cefSample, nil)
	}
}

func getCEFNormalizer() (INormalizer, error) {
	cfg := Config{
		Name: "test",
		Kind: KindCEF,
		Mapping: []MappingConfig{
			{SourceField: "deviceVendor", EventField: "UserString1"},
			{SourceField: "deviceProduct", EventField: "UserString2"},
			{SourceField: "deviceVersion", EventField: "UserString3"},
			{SourceField: "deviceEventClassId", EventField: "UserString4"},
			{SourceField: "name", EventField: "UserString5"},
			{SourceField: "severity", EventField: "UserString6"},
			{SourceField: "sourceAddress", EventField: "UserString7"},
			{SourceField: "deviceAction", EventField: "UserString8"},
			{SourceField: "destinationAddress", EventField: "DestinationIPAddress"},
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return New(cfg)
}
