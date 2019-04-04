package normalizers

import (
	"gotest.tools/assert"
	"testing"
)

var regexpSample = []byte(`<165>1 2003-10-11T22:14:15.003Z host.example.com logger - ID1714 test message`)
var regExpression = `<(?P<pri>\d+)>(?P<version>\d+) (?P<time>\d+-\d+-\d+T\d+:\d+:\d+\.\d{3}Z) (?P<host>\S+) (?P<app>\S+) \S+ (?P<mid>\S+) (?P<msg>.+)`

func TestRegexp(t *testing.T) {
	n, err := getRegexpNormalizer()
	assert.Equal(t, err, nil)
	event := n.Normalize(regexpSample, nil)
	assert.Equal(t, event.UserString1, "165")
	assert.Equal(t, event.UserString2, "1")
	assert.Equal(t, event.UserString3, "2003-10-11T22:14:15.003Z")
	assert.Equal(t, event.UserString4, "host.example.com")
	assert.Equal(t, event.UserString5, "logger")
	assert.Equal(t, event.UserString6, "ID1714")
	assert.Equal(t, event.UserString7, "test message")
}

func BenchmarkRegexp(b *testing.B) {
	n, _ := getRegexpNormalizer()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Normalize(regexpSample, nil)
	}
}

func getRegexpNormalizer() (INormalizer, error) {
	cfg := Config{
		Name:        "test",
		Kind:        KindRegexp,
		Expressions: []string{regExpression},
		Mapping: []MappingConfig{
			{SourceField: "pri", EventField: "UserString1"},
			{SourceField: "version", EventField: "UserString2"},
			{SourceField: "time", EventField: "UserString3"},
			{SourceField: "host", EventField: "UserString4"},
			{SourceField: "app", EventField: "UserString5"},
			{SourceField: "mid", EventField: "UserString6"},
			{SourceField: "msg", EventField: "UserString7"},
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return New(cfg)
}
