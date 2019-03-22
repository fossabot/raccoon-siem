package activeLists

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestExpirationTree(t *testing.T) {
	key := "testKey"
	tree := createExpirationTree()
	tree.touch(key, time.Now().UnixNano()+time.Millisecond.Nanoseconds())
	assert.Equal(t, len(tree.getExpiredKeys()), 0)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, len(tree.getExpiredKeys()), 1)
	tree.del(key)
	assert.Equal(t, len(tree.getExpiredKeys()), 0)
}

func BenchmarkExpirationTreeTouch(b *testing.B) {
	key := "testKey"
	tree := createExpirationTree()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.touch(key, time.Now().UnixNano()+time.Millisecond.Nanoseconds())
	}
}

func BenchmarkExpirationGetExpired(b *testing.B) {
	key := "testKey"
	tree := createExpirationTree()
	tree.touch(key, time.Now().UnixNano()+time.Millisecond.Nanoseconds())
	time.Sleep(2 * time.Millisecond)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.getExpiredKeys()
	}
}
