package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		if err := r.sendSlack(slackNotification{Text: n.message}); err != nil {
			fmt.Println(err)
		}
	}
}

func (r *Notifier) sendSlack(n slackNotification) error {
	body, err := json.Marshal(n)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, r.slackURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	_, err = r.httpClient.Do(req)
	return err
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
