package misc

import (
	"sync"
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int)

	go func() {
		for range ch {

		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- 1
	}
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}
