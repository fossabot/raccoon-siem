package normalization

import (
	"gotest.tools/assert"
	"testing"
)

func TestConverter(t *testing.T) {
	assert.Equal(t, StringToInt(" 032 "), int64(32))
	assert.Equal(t, StringToFloat(" 3.14 "), 3.14)
	assert.Equal(t, StringToBool(" true "), true)
	assert.Assert(t, StringToTime(" 2019-03-29T15:17:09.202Z ") != 0)
}

func BenchmarkStringToInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToInt(" 32 ")
	}
}

func BenchmarkStringToFloat(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToFloat(" 3.14 ")
	}
}

func BenchmarkStringToBool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToBool(" true ")
	}
}

func BenchmarkStringToTime(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToTime(" 2019-03-29T15:17:09.202Z ")
	}
}
