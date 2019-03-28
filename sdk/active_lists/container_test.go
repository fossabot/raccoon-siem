package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"log"
	"testing"
	"time"
)

var testContainer1, testContainer2 *Container

const testListName = "test"

func init() {
	var err error

	lists := []Config{{Name: testListName}}

	testContainer1, err = NewContainer(
		lists,
		"1",
		"nats://localhost:4222",
		"http://localhost:9200")
	if err != nil {
		log.Fatal(err)
	}

	testContainer2, err = NewContainer(
		lists,
		"2",
		"nats://localhost:4222",
		"http://localhost:9200")
	if err != nil {
		log.Fatal(err)
	}
}

func TestActiveList(t *testing.T) {
	bytesIn := int64(10000000)
	event := &normalization.Event{
		SourceIPAddress:      "192.168.2.2",
		DestinationIPAddress: "192.168.2.4",
		Message:              "attack",
		RequestBytesIn:       bytesIn,
	}
	keyFields := []string{"SourceIPAddress", "DestinationIPAddress"}

	// Set multiple fields:
	// 1. Check local result
	// 2. Check replicated result

	mapping := []Mapping{
		{EventField: "SourceIPAddress", ALField: "ip"},
		{EventField: "Message", ALField: "msg"},
		{EventField: "RequestBytesIn", ALField: "bytes_in"},
	}

	testContainer1.Set(testListName, keyFields, mapping, event)
	assert.Equal(t, normalization.ToInt64(testContainer1.Get(testListName, "bytes_in", keyFields, event)), bytesIn)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, normalization.ToInt64(testContainer2.Get(testListName, "bytes_in", keyFields, event)), bytesIn)

	// Update single field and add new one:
	// 1. Check local result
	// 2. Check the rest of fields didn't change
	// 3. Check replicated result

	event.Message = "attack2"
	event.RequestStatus = "200"
	mapping2 := []Mapping{
		{EventField: "Message", ALField: "msg"},
		{EventField: "RequestStatus", ALField: "status"},
	}

	testContainer2.Set(testListName, keyFields, mapping2, event)
	assert.Equal(t, testContainer2.Get(testListName, "status", keyFields, event), event.RequestStatus)
	assert.Equal(t, testContainer2.Get(testListName, "msg", keyFields, event), event.Message)
	assert.Equal(t, normalization.ToInt64(testContainer2.Get(testListName, "bytes_in", keyFields, event)), bytesIn)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, testContainer1.Get(testListName, "status", keyFields, event), event.RequestStatus)
	assert.Equal(t, testContainer1.Get(testListName, "msg", keyFields, event), event.Message)
	assert.Equal(t, normalization.ToInt64(testContainer1.Get(testListName, "bytes_in", keyFields, event)), bytesIn)

	// Delete key
	// 1. Check local result
	// 2. Check replicated result

	testContainer1.Del(testListName, keyFields, event)
	assert.Equal(t, testContainer1.Get(testListName, "status", keyFields, event), "")
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, testContainer2.Get(testListName, "status", keyFields, event), "")

	// Add one record
	// 1. Check it is read from storage by fresh container

	testContainer1.Set(testListName, keyFields, mapping, event)
	time.Sleep(2 * time.Second)
	testContainer3, err := NewContainer(
		[]Config{{Name: testListName}},
		"3",
		"nats://localhost:4222",
		"http://localhost:9200")
	assert.Equal(t, err, nil)
	assert.Equal(t, testContainer3.Get(testListName, "msg", keyFields, event), event.Message)
}

func BenchmarkALSet(b *testing.B) {
	bytesIn := int64(10000000)
	event := &normalization.Event{
		SourceIPAddress:      "192.168.2.2",
		DestinationIPAddress: "192.168.2.4",
		Message:              "attack",
		RequestBytesIn:       bytesIn,
	}

	mapping := []Mapping{
		{EventField: "SourceIPAddress", ALField: "ip"},
		{EventField: "Message", ALField: "msg"},
		{EventField: "RequestBytesIn", ALField: "bytes_in"},
	}

	keyFields := []string{"SourceIPAddress", "DestinationIPAddress"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testContainer1.Set(testListName, keyFields, mapping, event)
	}
}
