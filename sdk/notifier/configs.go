package notifier

type Config struct {
	Name     string `json:"name"`
	SlackURL string `json:"slackURL"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	return nil
}
