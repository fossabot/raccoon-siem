package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
)

var testListName = "test"
var testListTTL = 0
var testConstant = "constant"

func TestContainer(t *testing.T) {
	container, err := getContainer()
	assert.Equal(t, err, nil)

	event := getEvent()
	mapping, keys := getMappingAndKeys()
	container.Set(testListName, keys, mapping, event)

	value := container.Get(testListName, "msg", keys, event)
	assert.Equal(t, value, event.Message)

	value = container.Get(testListName, "severity", keys, event)
	assert.Equal(t, normalization.ToFieldType("Severity", value), event.Severity)

	value = container.Get(testListName, "float", keys, event)
	assert.Equal(t, normalization.ToFieldType("UserFloat1", value), event.UserFloat1)

	value = container.Get(testListName, "is_incident", keys, event)
	assert.Equal(t, normalization.ToFieldType("Incident", value), event.Incident)

	value = container.Get(testListName, "const_value", keys, event)
	assert.Equal(t, value, testConstant)

	container.Del(testListName, keys, event)
	value = container.Get(testListName, "msg", keys, event)
	assert.Equal(t, value, "")
}

func BenchmarkContainerSet(b *testing.B) {
	container, _ := getContainer()
	event := getEvent()
	mapping, keys := getMappingAndKeys()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container.Set(testListName, keys, mapping, event)
	}
}

func BenchmarkContainerGet(b *testing.B) {
	container, _ := getContainer()
	event := getEvent()
	_, keys := getMappingAndKeys()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container.Get(testListName, "msg", keys, event)
	}
}

func getContainer() (*Container, error) {
	return NewContainer([]Config{{
		Name: testListName,
		TTL:  int64(testListTTL),
	}}, "localhost:6379")
}

func getEvent() *normalization.Event {
	return &normalization.Event{
		Incident:             true,
		Message:              "test incident",
		Severity:             3,
		UserFloat1:           3.14,
		UserFloat1Label:      "pi",
		DestinationIPAddress: "192.168.1.1",
		DestinationPort:      "53",
	}
}

func getMappingAndKeys() (mapping []Mapping, keyFields []string) {
	mapping = []Mapping{
		{EventField: "Message", ALField: "msg"},
		{EventField: "Severity", ALField: "severity"},
		{EventField: "DestinationIPAddress", ALField: "ip"},
		{EventField: "UserFloat1", ALField: "float"},
		{EventField: "Incident", ALField: "is_incident"},
		{Constant: "constant", ALField: "const_value"},
	}
	keyFields = []string{"DestinationIPAddress", "DestinationPort"}
	return
}
