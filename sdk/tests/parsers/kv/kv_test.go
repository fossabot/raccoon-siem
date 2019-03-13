package kv

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/parsers/kv"
	"gotest.tools/assert"
	"testing"
)

var benchmarkInput = []byte("key1=value1;key2=value2;key3=value3;key4=value4;key5=value5;key6=value6;key7=value7;key8=value8;")
var semicolonEqualInput = []byte("key1=value1;key2=value2;")
var colonEqualInput = []byte("key1=value1:key2=value2;")
var minusCommaInput = []byte("key1,value1-key2,value2")
var spacedInput = []byte(" first key = value1, second key = value2 ")

const (
	equal     = byte('=')
	semicolon = byte(';')
	comma     = byte(',')
	colon     = byte(':')
	minus     = byte('-')
)

func BenchmarkParser(b *testing.B) {
	b.ReportAllocs()
	parser, _ := kv.NewParser(semicolon, equal)

	for i := 0; i < b.N; i++ {
		parser.Parse(benchmarkInput)
	}
}

func TestParser(t *testing.T) {
	var res map[string]string
	var success bool

	p, err := kv.NewParser(semicolon, semicolon)
	assert.Error(t, err, "kv and pair separators must be different")

	p, err = kv.NewParser(semicolon, equal)
	res, success = p.Parse(semicolonEqualInput)
	assert.Assert(t, success, true)
	assert.Equal(t, res["key1"], "value1")
	assert.Equal(t, res["key2"], "value2")

	p, err = kv.NewParser(colon, equal)
	res, success = p.Parse(colonEqualInput)
	assert.Assert(t, success, true)
	assert.Equal(t, res["key1"], "value1")
	assert.Equal(t, res["key2"], "value2;")

	p, err = kv.NewParser(minus, comma)
	res, success = p.Parse(minusCommaInput)
	assert.Assert(t, success, true)
	assert.Equal(t, res["key1"], "value1")
	assert.Equal(t, res["key2"], "value2")

	p, err = kv.NewParser(comma, equal)
	res, success = p.Parse(spacedInput)
	assert.Assert(t, success, true)
	assert.Equal(t, res["first key"], "value1", fmt.Sprintf("Failed: %s != %s", res["first key"], "value1"))
	assert.Equal(t, res["second key"], "value2")
}
