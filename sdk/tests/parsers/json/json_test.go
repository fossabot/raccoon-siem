package json

import (
	jp "github.com/tephrocactus/raccoon-siem/sdk/parsers/json"
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
	parser, _ := jp.NewParser(jp.Config{})
	result, success := parser.Parse(sample)
	assert.Equal(t, success, true)
	assert.Equal(t, result["first"], "timestamp")
	assert.Equal(t, result["timestamp"], "2018-05-24 23:15:07")
	assert.Equal(t, result["connection_id"], "12")
	assert.Equal(t, result["account.user"], "user")
	assert.Equal(t, result["login.org"], "null")
	assert.Equal(t, result["connection_data.connection_type"], "tcp/ip")
	assert.Equal(t, result["level"], "-1")

	_, success = parser.Parse(sampleWithoutOpenCurlyBracket)
	assert.Equal(t, success, false)

	_, success = parser.Parse(sampleWithoutCloseCurlyBracket)
	assert.Equal(t, success, false)
}

func BenchmarkJSONParser(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parser, _ := jp.NewParser(jp.Config{})
		parser.Parse(sample)
	}
}
