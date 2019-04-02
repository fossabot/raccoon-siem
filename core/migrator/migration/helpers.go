package migration

import (
	"fmt"
	"time"
)

// Сообщает о результатах миграции
func ReportAppliedFiles(files []string, processingBegan time.Time) {
	fmt.Printf("[*] %d migrations applied in %s\n", len(files), time.Since(processingBegan))
	if len(files) > 0 {
		for _, file := range files {
			fmt.Printf("    - %s\n", file)
		}
	}
}
