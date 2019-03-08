package json

import (
	jp "github.com/tephrocactus/raccoon-siem/sdk/parsers/json"
	"github.com/tidwall/gjson"
	"gotest.tools/assert"
	"testing"
)

var sample = []byte(`
{
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
}
`)

var paths = []string{
	"timestamp",
	"account.user",
	"account.group",
	"login.org",
	"connection_data.status",
	"account.host",
	"not.exists",
	"string_arr",
	"медведи и балалайки",
	"level",
}

var pathsByte = [][]byte{
	[]byte("timestamp"),
	[]byte("account.user"),
	[]byte("account.group"),
	[]byte("login.org"),
	[]byte("connection_data.status"),
	[]byte("account.host"),
	[]byte("not.exists"),
	[]byte("string_arr"),
	[]byte("медведи и балалайки"),
	[]byte("level"),
}

func BenchmarkGJSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			_ = gjson.GetBytes(sample, path).String()
		}
	}
}

func TestJSONParser(t *testing.T) {
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[0]))), "2018-05-24 23:15:07")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[1]))), "user")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[2]))), "")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[3]))), "null")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[4]))), "0")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[5]))), `localh\"ost`)
	assert.Assert(t, jp.GetValue(sample, []byte(paths[6])) == nil)
	assert.Assert(t, jp.GetValue(sample, []byte(paths[7])) == nil)
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[8]))), "true")
	assert.Equal(t, string(jp.GetValue(sample, []byte(paths[9]))), "-1")
}

func BenchmarkJSONParser(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, path := range pathsByte {
			jp.GetValue(sample, path)
		}
	}
}
