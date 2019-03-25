package regexp

import (
	"gotest.tools/assert"
	"regexp"
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
	exps := compileExpressions()

	result, ok := Parse(messages[0], exps)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["pri"], []byte("165"))
	assert.DeepEqual(t, result["version"], []byte("1"))
	assert.DeepEqual(t, result["time"], []byte("2003-10-11T22:14:15.003Z"))
	assert.DeepEqual(t, result["host"], []byte("host.example.com"))
	assert.DeepEqual(t, result["app"], []byte("logger"))
	assert.DeepEqual(t, result["mid"], []byte("ID1714"))
	assert.DeepEqual(t, result["msg"], []byte("test message"))

	result, ok = Parse(messages[1], exps)
	assert.DeepEqual(t, ok, true)
	assert.DeepEqual(t, result["pri"], []byte("34"))
	assert.DeepEqual(t, result["time"], []byte("Oct 11 22:14:15"))
	assert.DeepEqual(t, result["host"], []byte("mymachine"))
	assert.DeepEqual(t, result["app"], []byte("su"))
	assert.DeepEqual(t, result["msg"], []byte("'su root' failed for lonvick on /dev/pts/8"))
}

func BenchmarkRegexp(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	exps := compileExpressions()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Parse(messages[0], exps)
	}
}

func compileExpressions() (exps []*regexp.Regexp) {
	for _, e := range expressions {
		exps = append(exps, regexp.MustCompile(e))
	}
	return
}
