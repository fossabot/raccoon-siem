package regexp

import (
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/regexp"
	"gotest.tools/assert"
	"testing"
)

var messages = [][]byte{
	[]byte(`<165>1 2003-10-11T22:14:15.003Z host.example.com logger - ID1714 test message`),
	[]byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`),
}

var expressions = []string{
	`<(?P<pri>\d+)>(?P<time>\S+ \d+ \d+:\d+:\d+) (?P<host>\S+) (?P<app>\S+): (?P<msg>.+)`,
	`<(?P<pri>\d+)>(?P<version>\d+) (?P<time>\d+-\d+-\d+T\d+:\d+:\d+\.\d{3}Z) (?P<host>\S+) (?P<app>\S+) \S+ (?P<mid>\S+) (?P<msg>.+)`,
}

func TestRegexp(t *testing.T) {
	_, err := regexp.NewParser("zero_groups", []string{`\S+`})
	assert.Assert(t, err != nil)

	_, err = regexp.NewParser("no_names", []string{`(\S+)`})
	assert.Assert(t, err != nil)

	p, err := regexp.NewParser("valid", expressions)
	assert.Assert(t, err == nil)

	result, ok := p.Parse(messages[0])
	assert.Equal(t, ok, true)
	assert.Equal(t, result["pri"], "165")
	assert.Equal(t, result["version"], "1")
	assert.Equal(t, result["time"], "2003-10-11T22:14:15.003Z")
	assert.Equal(t, result["host"], "host.example.com")
	assert.Equal(t, result["app"], "logger")
	assert.Equal(t, result["mid"], "ID1714")
	assert.Equal(t, result["msg"], "test message")

	result, ok = p.Parse(messages[1])
	assert.Equal(t, ok, true)
	assert.Equal(t, result["pri"], "34")
	assert.Equal(t, result["time"], "Oct 11 22:14:15")
	assert.Equal(t, result["host"], "mymachine")
	assert.Equal(t, result["app"], "su")
	assert.Equal(t, result["msg"], "'su root' failed for lonvick on /dev/pts/8")
}

func BenchmarkRegexp(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	p, _ := regexp.NewParser("regexp", expressions)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(messages[0])
	}
}
