package cef

import (
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/cef"
	"gotest.tools/assert"
	"testing"
)

var sample = []byte(`CEF:0|енот|threatmanager|1.0|100|detected a \| in message|10|src=10.0.0.1 act=blocked message dst=1.1.1.1`)

func TestCEF(t *testing.T) {
	p := cef.Parser{}

	_, ok := p.Parse([]byte("invalid"))
	assert.Equal(t, ok, false)

	result, ok := p.Parse(sample)
	assert.Equal(t, ok, true)
	assert.Equal(t, result["deviceVendor"], "енот")
	assert.Equal(t, result["deviceProduct"], "threatmanager")
	assert.Equal(t, result["deviceVersion"], "1.0")
	assert.Equal(t, result["deviceEventClassId"], "100")
	assert.Equal(t, result["name"], `detected a \| in message`)
	assert.Equal(t, result["severity"], "10")
	assert.Equal(t, result["sourceAddress"], "10.0.0.1")
	assert.Equal(t, result["deviceAction"], "blocked message")
	assert.Equal(t, result["destinationAddress"], "1.1.1.1")
}

func BenchmarkCEF(b *testing.B) {
	b.ReportAllocs()
	p := cef.Parser{}
	for i := 0; i < b.N; i++ {
		p.Parse(sample)
	}
}
