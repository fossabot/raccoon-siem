package rfc3164

import (
	"gotest.tools/assert"
	"testing"
)

var sampleRFC3164 = []byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`)

func TestSyslogRFC3164(t *testing.T) {
	result := make(map[string][]byte)
	callback := func(key string, value []byte) {
		result[key] = value
	}

	ok := Parse(sampleRFC3164, callback)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["facility"], []byte("4"))
	assert.DeepEqual(t, result["severity"], []byte("2"))
	assert.DeepEqual(t, result["time"], []byte("Oct 11 22:14:15"))
	assert.DeepEqual(t, result["host"], []byte("mymachine"))
	assert.DeepEqual(t, result["app"], []byte("su"))
	assert.DeepEqual(t, result["pid"], []byte(nil))
	assert.DeepEqual(t, result["mid"], []byte(nil))
	assert.DeepEqual(t, result["msg"], []byte("'su root' failed for lonvick on /dev/pts/8"))
}

func BenchmarkSyslogRFC3164(b *testing.B) {
	callback := func(key string, value []byte) {}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(sampleRFC3164, callback)
	}
}
