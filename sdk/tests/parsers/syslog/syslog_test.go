package syslog

import (
	"github.com/tephrocactus/raccoon-siem/sdk/parsers"
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/syslog/rfc3164"
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/syslog/rfc5424"
	"gotest.tools/assert"
	"testing"
)

var sampleRFC5424 = []byte(`<165>1 2003-10-11T22:14:15.003Z host.example.com logger - ID1714 [SID service="rest" version="1.0.1" status="500"] Internal Server Error`)
var sampleRFC3164 = []byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`)

func TestSyslogRFC5424(t *testing.T) {
	p, _ := rfc5424.NewParser(rfc5424.Config{
		BaseConfig: parsers.BaseConfig{Name: "test"},
	})

	result, ok := p.Parse(sampleRFC5424)
	assert.Equal(t, ok, true)
	assert.Equal(t, result["facility"], "20")
	assert.Equal(t, result["severity"], "5")
	assert.Equal(t, result["time"], "2003-10-11T22:14:15.003Z")
	assert.Equal(t, result["host"], "host.example.com")
	assert.Equal(t, result["app"], "logger")
	assert.Equal(t, result["pid"], "-")
	assert.Equal(t, result["mid"], "ID1714")
	assert.Equal(t, result["service"], "rest")
	assert.Equal(t, result["version"], "1.0.1")
	assert.Equal(t, result["status"], "500")
	assert.Equal(t, result["msg"], "Internal Server Error")

	_, ok = p.Parse(sampleRFC3164)
	assert.Equal(t, ok, false)
}

func TestSyslogRFC3164(t *testing.T) {
	p, _ := rfc3164.NewParser(rfc3164.Config{
		BaseConfig: parsers.BaseConfig{Name: "test"},
	})

	result, ok := p.Parse(sampleRFC3164)
	assert.Equal(t, ok, true)
	assert.Equal(t, result["facility"], "4")
	assert.Equal(t, result["severity"], "2")
	assert.Equal(t, result["time"], "Oct 11 22:14:15")
	assert.Equal(t, result["host"], "mymachine")
	assert.Equal(t, result["app"], "su")
	assert.Equal(t, result["pid"], "")
	assert.Equal(t, result["mid"], "")
	assert.Equal(t, result["msg"], "'su root' failed for lonvick on /dev/pts/8")
}

func BenchmarkSyslogRFC5424(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	p, _ := rfc5424.NewParser(rfc5424.Config{
		BaseConfig: parsers.BaseConfig{Name: "test"},
	})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(sampleRFC5424)
	}
}

func BenchmarkSyslogRFC3164(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	p, _ := rfc3164.NewParser(rfc3164.Config{
		BaseConfig: parsers.BaseConfig{Name: "test"},
	})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(sampleRFC3164)
	}
}
