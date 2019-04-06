package csv

import (
	"gotest.tools/assert"
	"testing"
)

var unquotedSample = []byte(`CEF:0,енот,threatmanager,1.0,100,detected,10,10.0.0.1,blocked message,1.1.1.1`)
var quotedSample = []byte(`"value1","value,2","value3"`)

func TestCSV(t *testing.T) {
	result := make(map[string][]byte)
	callback := func(key string, value []byte) {
		result[key] = value
	}

	ok := Parse(unquotedSample, ',', callback)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["0"], []byte("CEF:0"))
	assert.DeepEqual(t, result["1"], []byte("енот"))
	assert.DeepEqual(t, result["2"], []byte("threatmanager"))
	assert.DeepEqual(t, result["3"], []byte("1.0"))
	assert.DeepEqual(t, result["4"], []byte("100"))
	assert.DeepEqual(t, result["5"], []byte(`detected`))
	assert.DeepEqual(t, result["6"], []byte("10"))
	assert.DeepEqual(t, result["7"], []byte("10.0.0.1"))
	assert.DeepEqual(t, result["8"], []byte("blocked message"))
	assert.DeepEqual(t, result["9"], []byte("1.1.1.1"))

	ok = Parse(quotedSample, ',', callback)
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, result["0"], []byte("value1"))
	assert.DeepEqual(t, result["1"], []byte("value,2"))
	assert.DeepEqual(t, result["2"], []byte("value3"))
}

func BenchmarkCEF(b *testing.B) {
	cb := func(key string, value []byte) {}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(unquotedSample, ',', cb)
	}
}
