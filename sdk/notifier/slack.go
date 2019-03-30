package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type slackNotification struct {
	Text string `json:"text"`
}

func (r *slackNotification) send(cli *http.Client, url string) error {
	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	_, err = cli.Do(req)
	return err
}
