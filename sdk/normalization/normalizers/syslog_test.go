package normalizers

import (
	"gotest.tools/assert"
	"testing"
)

var syslogSample = []byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`)

func TestSyslog(t *testing.T) {
	n, err := getSyslogNormalizer()
	assert.Equal(t, err, nil)
	event := n.Normalize(syslogSample, nil)
	assert.Assert(t, event.UserString1 != "")
	assert.Assert(t, event.UserString2 != "")
	assert.Equal(t, event.UserString3, "Oct 11 22:14:15")
	assert.Equal(t, event.UserString4, "mymachine")
	assert.Equal(t, event.UserString5, "su")
	assert.Equal(t, event.UserString8, `'su root' failed for lonvick on /dev/pts/8`)
}

func BenchmarkSyslog(b *testing.B) {
	n, _ := getSyslogNormalizer()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Normalize(syslogSample, nil)
	}
}

func getSyslogNormalizer() (INormalizer, error) {
	cfg := Config{
		Name: "test",
		Kind: KindSyslog,
		Mapping: []MappingConfig{
			{SourceField: "facility", EventField: "UserString1"},
			{SourceField: "severity", EventField: "UserString2"},
			{SourceField: "time", EventField: "UserString3"},
			{SourceField: "host", EventField: "UserString4"},
			{SourceField: "app", EventField: "UserString5"},
			{SourceField: "pid", EventField: "UserString6"},
			{SourceField: "mid", EventField: "UserString7"},
			{SourceField: "msg", EventField: "UserString8"},
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return New(cfg)
}
