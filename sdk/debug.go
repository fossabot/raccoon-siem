package sdk

import (
	"fmt"
	"runtime"
	"time"
)

var Debug bool

func PrintMemUsage(interval time.Duration) {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		fmt.Printf("[memory] Current: %v M \tMaximum: %v M\n", bToMb(m.Alloc), bToMb(m.Sys))
		time.Sleep(interval)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
