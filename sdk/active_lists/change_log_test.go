package activeLists

import (
	"gopkg.in/vmihailenco/msgpack.v4"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestChangeLogEncodeDecode(t *testing.T) {
	chLog := testGetChangeLog()

	encoded, err := msgpack.Marshal(&chLog)
	assert.Equal(t, err, nil)

	var decoded changeLog
	err = msgpack.Unmarshal(encoded, &decoded)
	assert.Equal(t, err, nil)

	assert.Equal(t, decoded.CID, chLog.CID)
	assert.Equal(t, decoded.ALName, chLog.ALName)
	assert.Equal(t, decoded.Op, chLog.Op)
	assert.Equal(t, decoded.Key, chLog.Key)
	assert.Equal(t, decoded.Version, chLog.Version)
	assert.DeepEqual(t, decoded.Record, chLog.Record)
}

func BenchmarkChangeLogEncode(b *testing.B) {
	chLog := testGetChangeLog()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(&chLog)
	}
}

func BenchmarkChangeLogDecode(b *testing.B) {
	chLog := testGetChangeLog()
	serialized, _ := msgpack.Marshal(&chLog)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deserialized := changeLog{}
		_ = msgpack.Unmarshal(serialized, &deserialized)
	}
}

func testGetChangeLog() changeLog {
	ts := time.Now().UnixNano()
	return changeLog{
		CID:     "testComponent",
		ALName:  "test",
		Op:      OpSet,
		Key:     "testKey",
		Version: ts,
		Record: Record{
			Version: ts,
			Fields:  map[string]string{"testField": "123"},
		},
	}
}
