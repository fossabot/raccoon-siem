package cef

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/cef"
	"gotest.tools/assert"
	"testing"
)

var sampleNoExt = []byte(`CEF:0|енот|threatmanager|1.0|100|detected a pipe in message|10|`)
var sampleExt = []byte(`CEF:0|енот|threatmanager|1.0|100|detected a pipe in message|10|src=10.0.0.1 act=blocked dst=1.1.1.1`)

func TestCEF(t *testing.T) {
	p := cef.Parser{}

	result, ok := p.Parse(sampleExt)
	assert.Equal(t, ok, true)
	fmt.Println(result)

	//result, ok = p.Parse(sampleNoExt)
	//assert.Equal(t, ok, true)
	//fmt.Println(result)

	//assert.Equal(t, result["facility"], "20")
	//assert.Equal(t, result["severity"], "5")
	//assert.Equal(t, result["time"], "2003-10-11T22:14:15.003Z")
	//assert.Equal(t, result["host"], "host.example.com")
	//assert.Equal(t, result["app"], "logger")
	//assert.Equal(t, result["pid"], "-")
	//assert.Equal(t, result["mid"], "ID1714")
	//assert.Equal(t, result["service"], "rest")
	//assert.Equal(t, result["version"], "1.0.1")
	//assert.Equal(t, result["status"], "500")
	//assert.Equal(t, result["msg"], "Internal Server Error")
}

func BenchmarkCEF(b *testing.B) {
	b.ReportAllocs()
	p := cef.Parser{}
	for i := 0; i < b.N; i++ {
		p.Parse(sampleExt)
	}
}
