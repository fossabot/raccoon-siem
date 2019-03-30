package normalization

import (
	"github.com/satori/go.uuid"
	"gotest.tools/assert"
	"log"
	"testing"
	"time"
)

func TestMsgPackEncodeDecode(t *testing.T) {
	event := testGetEvent()
	encodedEvent, err := event.ToMsgPack()
	assert.Equal(t, err, nil)
	decodedEvent := new(Event)
	err = decodedEvent.FromMsgPack(encodedEvent)
	assert.Equal(t, err, nil)
	assert.DeepEqual(t, decodedEvent, event)
}

func BenchmarkMsgPackEncode(b *testing.B) {
	event := testGetEvent()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = event.ToMsgPack()
	}
}

func BenchmarkMsgPackDecode(b *testing.B) {
	event := testGetEvent()
	encodedEvent, _ := event.ToMsgPack()
	decodedEvent := new(Event)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = decodedEvent.FromMsgPack(encodedEvent)
	}
}

func TestJSONEncodeDecode(t *testing.T) {
	event := testGetEvent()
	encodedEvent, err := event.ToJSON()
	assert.Equal(t, err, nil)
	decodedEvent := new(Event)
	err = decodedEvent.FromJSON(encodedEvent)
	assert.Equal(t, err, nil)
	assert.DeepEqual(t, decodedEvent, event)
}

func BenchmarkJSONEncode(b *testing.B) {
	event := testGetEvent()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := event.ToJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkJSONDecode(b *testing.B) {
	event := testGetEvent()
	encodedEvent, _ := event.ToJSON()
	decodedEvent := new(Event)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := decodedEvent.FromJSON(encodedEvent); err != nil {
			log.Fatal(err)
		}
	}
}

func testGetEvent() *Event {
	return &Event{
		ID:                   uuid.NewV4().String(),
		Tag:                  uuid.NewV4().String(),
		Timestamp:            time.Now().UnixNano() / 1000,
		BaseEventCount:       1,
		AggregatedEventCount: 1,
		AggregationRuleName:  uuid.NewV4().String(),
		CollectorIPAddress:   uuid.NewV4().String(),
		CollectorDNSName:     uuid.NewV4().String(),
		CorrelationRuleName:  uuid.NewV4().String(),
		CorrelatorDNSName:    uuid.NewV4().String(),
		CorrelatorIPAddress:  uuid.NewV4().String(),
		BaseEventIDs:         strSlice{uuid.NewV4().String(), uuid.NewV4().String(), uuid.NewV4().String()},
		SourceID:             uuid.NewV4().String(),
		Message:              uuid.NewV4().String(),
		OriginTimestamp:      time.Now().UnixNano() / 1000,
		Severity:             1,
		UserFloat1:           1.11,
		Incident:             true,
	}
}
