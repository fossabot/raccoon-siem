package json

import (
	"gotest.tools/assert"
	"testing"
)

var sampleWithoutOpenCurlyBracket = []byte(`"key": 2}`)
var sampleWithoutCloseCurlyBracket = []byte(`{"key": 2`)

var sample = []byte(`{
  "first": "timestamp",

  "timestamp": "2018-05-24 23:15:07",

  "id": 0,

  "class": "connection",

  "event": "connect",

  "connection_id": 12,

  "медведи и балалайки": true,

  "account": {
    "user": "user",
    "host": "localh\"ost",
    "domain": null,
    "login": "undefined",
    "group": ""
  },

  "string_arr": [ "1", "2", "3[" ],

  "obj_arr": [ { "k": null }, { "k": "s}tr", "v": [ "a", "b", "c" ] } ],

  "login": {
    "user": "user",
    "os": "",
    "ip": "::1",
    "proxy": "",
    "org": null
  },

  "connection_data": {
    "connection_type": "tcp/ip",
    "status": 0,
    "db": "bank_db"
  },

  "level": -1
}`)

func TestJSONParser(t *testing.T) {
	var result map[string][]byte
	var success bool
	for i := 0; i < 5; i++ {
		result, success = Parse(sample)
	}
	assert.Equal(t, success, true)

	assert.DeepEqual(t, result["first"], []byte("timestamp"))
	assert.DeepEqual(t, result["timestamp"], []byte("2018-05-24 23:15:07"))
	assert.DeepEqual(t, result["connection_id"], []byte("12"))
	assert.DeepEqual(t, result["account.user"], []byte("user"))
	assert.DeepEqual(t, result["login.org"], []byte("null"))
	assert.DeepEqual(t, result["connection_data.connection_type"], []byte("tcp/ip"))
	assert.DeepEqual(t, result["level"], []byte("-1"))

	_, success = Parse(sampleWithoutOpenCurlyBracket)
	assert.Equal(t, success, false)

	_, success = Parse(sampleWithoutCloseCurlyBracket)
	assert.Equal(t, success, false)
}

func BenchmarkJSONParser(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse(sample)
	}
}
