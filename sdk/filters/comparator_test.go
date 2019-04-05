package filters

import (
	"gotest.tools/assert"
	"testing"
)

func TestInSubnet(t *testing.T) {
	comp := comparator{}
	assert.Equal(t, comp.inSubnet("10.10.10.2", "10.10.10.0/24"), true)
	assert.Equal(t, comp.inSubnet("10.10.14.2", "10.10.10.0/24"), false)
}

func TestContains(t *testing.T) {
	comp := comparator{}
	assert.Equal(t, comp.contains("123needle321", "needle"), true)
	assert.Equal(t, comp.contains("123needle321", "stack"), false)
}

func BenchmarkInSubnet(b *testing.B) {
	comp := comparator{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp.inSubnet("10.10.10.2", "10.10.10.0/24")
	}
}
