package notifier

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type notification struct {
	message string
	url     string
}

type Notifier struct {
	httpClient *http.Client
	slackURL   string
	maxRetries int
	queue      chan notification
}

func (r *Notifier) Notify(msg, url string) {
	r.queue <- notification{message: msg, url: url}
}

func (r *Notifier) worker() {
	for n := range r.queue {
		if r.slackURL != "" {
			slack := slackNotification{Text: fmt.Sprintf("%s: <%s>", n.message, n.url)}
			if err := slack.send(r.httpClient, r.slackURL); err != nil {
				log.Println(err)
			}
		}
	}
}

func New(cfg Config) (*Notifier, error) {
	n := &Notifier{
		queue:      make(chan notification, 4096),
		httpClient: &http.Client{Timeout: 5 * time.Second},
		maxRetries: 3,
		slackURL:   cfg.SlackURL,
	}

	go n.worker()
	return n, nil
}
