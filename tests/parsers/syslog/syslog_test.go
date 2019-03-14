package syslog

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers/parsers/syslog/rfc3164"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers/parsers/syslog/rfc5424"
	"gotest.tools/assert"
	"testing"
)

var sampleRFC5424 = []byte(`<165>1 2003-10-11T22:14:15.003Z host.example.com logger - ID1714 [SID service="rest" version="1.0.1" status="500"] Internal Server Error`)
var sampleRFC3164 = []byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`)

func TestSyslogRFC5424(t *testing.T) {
	result, ok := rfc5424.Parse(sampleRFC5424)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["facility"], []byte("20"))
	assert.DeepEqual(t, result["severity"], []byte("5"))
	assert.DeepEqual(t, result["time"], []byte("2003-10-11T22:14:15.003Z"))
	assert.DeepEqual(t, result["host"], []byte("host.example.com"))
	assert.DeepEqual(t, result["app"], []byte("logger"))
	assert.DeepEqual(t, result["pid"], []byte("-"))
	assert.DeepEqual(t, result["mid"], []byte("ID1714"))
	assert.DeepEqual(t, result["service"], []byte("rest"))
	assert.DeepEqual(t, result["version"], []byte("1.0.1"))
	assert.DeepEqual(t, result["status"], []byte("500"))
	assert.DeepEqual(t, result["msg"], []byte("Internal Server Error"))

	_, ok = rfc5424.Parse(sampleRFC3164)
	assert.Equal(t, ok, false)
}

func TestSyslogRFC3164(t *testing.T) {
	result, ok := rfc3164.Parse(sampleRFC3164)
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

func BenchmarkSyslogRFC5424(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rfc5424.Parse(sampleRFC5424)
	}
}

func BenchmarkSyslogRFC3164(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rfc3164.Parse(sampleRFC3164)
	}
}
