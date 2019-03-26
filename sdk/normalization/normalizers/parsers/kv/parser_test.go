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
	var res map[string][]byte
	var success bool

	res, success = Parse(semicolonEqualInput, semicolon, equal)
	assert.DeepEqual(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	res, success = Parse(colonEqualInput, colon, equal)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2;"))

	res, success = Parse(minusCommaInput, minus, comma)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	res, success = Parse(spacedInput, comma, equal)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["first key"], []byte("value1"))
	assert.DeepEqual(t, res["second key"], []byte("value2"))

	res, success = Parse(escapedInput, comma, equal)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1\\=key2"], []byte("value1\\,value2"))
	assert.DeepEqual(t, res["second key"], []byte("value2"))

	res, success = Parse(commonInput, space, equal)
	assert.Assert(t, success, true)
	assert.DeepEqual(t, res["key1"], []byte("value1"))
	assert.DeepEqual(t, res["key2"], []byte("value2"))

	res, success = Parse([]byte(""), space, equal)
	assert.Assert(t, success, false)
}

func BenchmarkParser(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse(benchmarkInput, semicolon, equal)
	}
}
