package kv

import (
	"gotest.tools/assert"
	"testing"
)

var benchmarkInput = []byte("key1=value1;key2=value2;key3=value3;key4=value4;key5=value5;key6=value6;key7=value7;key8=value8;")
var semicolonEqualInput = []byte("key1=value1;key2=value2;")
var colonEqualInput = []byte("key1=value1:key2=value2;")
var minusCommaInput = []byte("key1,value1-key2,value2")
var spacedInput = []byte(" first key = value1, second key = value2 ")
var escapedInput = []byte("key1\\=key2=value1\\,value2,second key = value2 ")
var commonInput = []byte("key1=value1 key2=value2")

const (
	equal     = byte('=')
	semicolon = byte(';')
	comma     = byte(',')
	colon     = byte(':')
	minus     = byte('-')
)

func TestParser(t *testing.T) {
	res := make(map[string][]byte)
	success := false
	callback := func(key string, value []byte) {
		res[key] = value
	}

	success = Parse(semicolonEqualInput, semicolon, equal, callback)
	assert.DeepEqual(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	success = Parse(colonEqualInput, colon, equal, callback)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2;"))

	success = Parse(minusCommaInput, minus, comma, callback)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	success = Parse(spacedInput, comma, equal, callback)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["first key"], []byte("value1"))
	assert.DeepEqual(t, res["second key"], []byte("value2"))

	success = Parse(escapedInput, comma, equal, callback)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1\\=key2"], []byte("value1\\,value2"))
	assert.DeepEqual(t, res["second key"], []byte("value2"))

	success = Parse(commonInput, space, equal, callback)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	success = Parse([]byte(""), space, equal, callback)
	assert.Assert(t, success, false)
}

func BenchmarkParser(b *testing.B) {
	cb := func(key string, value []byte) {}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(benchmarkInput, semicolon, equal, cb)
	}
}
