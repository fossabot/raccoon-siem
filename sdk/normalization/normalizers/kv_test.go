package normalizers

import (
	"gotest.tools/assert"
	"testing"
)

var kvSample = []byte("key1=value1;key2=value2;key3=value3;key4=value4;key5=value5;key6=value6;key7=value7;key8=value8;")

func TestKV(t *testing.T) {
	n, err := getKVNormalizer()
	assert.Equal(t, err, nil)
	event := n.Normalize(kvSample, nil)
	assert.Equal(t, event.UserString1, "value1")
	assert.Equal(t, event.UserString2, "value2")
	assert.Equal(t, event.UserString3, "value3")
	assert.Equal(t, event.UserString4, "value4")
	assert.Equal(t, event.UserString5, "value5")
	assert.Equal(t, event.UserString6, "value6")
	assert.Equal(t, event.UserString7, "value7")
	assert.Equal(t, event.UserString8, "value8")
}

func BenchmarkKV(b *testing.B) {
	n, _ := getKVNormalizer()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Normalize(kvSample, nil)
	}
}

func getKVNormalizer() (INormalizer, error) {
	cfg := Config{
		Name:          "test",
		Kind:          KindKV,
		KVDelimiter:   "=",
		PairDelimiter: ";",
		Mapping: []MappingConfig{
			{SourceField: "key1", EventField: "UserString1"},
			{SourceField: "key2", EventField: "UserString2"},
			{SourceField: "key3", EventField: "UserString3"},
			{SourceField: "key4", EventField: "UserString4"},
			{SourceField: "key5", EventField: "UserString5"},
			{SourceField: "key6", EventField: "UserString6"},
			{SourceField: "key7", EventField: "UserString7"},
			{SourceField: "key8", EventField: "UserString8"},
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return New(cfg)
}
