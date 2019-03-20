package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	al := New()
	event := &normalization.Event{
		SourceIPAddress:      "192.168.2.2",
		DestinationIPAddress: "192.168.2.4",
	}
	fields := []string{"SourceIPAddress", "DestinationIPAddress"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		al.Set("dos", fields, event)
	}
}
