package cef

import (
	"gotest.tools/assert"
	"testing"
)

var sample = []byte(`CEF:0|енот|threatmanager|1.0|100|detected a \| in message|10|src=10.0.0.1 act=blocked message dst=1.1.1.1`)

func TestCEF(t *testing.T) {
	result, ok := Parse([]byte("invalid"))
	assert.Equal(t, ok, false)

	result, ok = Parse(sample)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["deviceVendor"], []byte("енот"))
	assert.DeepEqual(t, result["deviceProduct"], []byte("threatmanager"))
	assert.DeepEqual(t, result["deviceVersion"], []byte("1.0"))
	assert.DeepEqual(t, result["deviceEventClassId"], []byte("100"))
	assert.DeepEqual(t, result["name"], []byte(`detected a \| in message`))
	assert.DeepEqual(t, result["severity"], []byte("10"))
	assert.DeepEqual(t, result["sourceAddress"], []byte("10.0.0.1"))
	assert.DeepEqual(t, result["deviceAction"], []byte("blocked message"))
	assert.DeepEqual(t, result["destinationAddress"], []byte("1.1.1.1"))
}

func BenchmarkCEF(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse(sample)
	}
}
