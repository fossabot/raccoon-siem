package notifier

import (
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestNotifier(t *testing.T) {
	n, err := New(Config{
		SlackURL: "https://hooks.slack.com/services/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",
	})
	assert.Equal(t, err, nil)
	n.Notify("test message", "")
	time.Sleep(2 * time.Second)
}
