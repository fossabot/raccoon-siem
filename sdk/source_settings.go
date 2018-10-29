package sdk

type SourceSettings struct {
	Name        string `yaml:"name,omitempty"`
	Kind        string `yaml:"kind,omitempty"`
	URL         string `yaml:"url,omitempty"`
	Channel     string `yaml:"channel,omitempty"`
	Queue       string `yaml:"queue,omitempty"`
	FilePath    string `yaml:"filePath,omitempty"`
	LoadBalance bool   `yaml:"loadBalance,omitempty"`
	Mode        string `yaml:"mode,omitempty"`
	Buffer      int    `yaml:"buffer,omitempty"`
}

func (s *SourceSettings) ID() string {
	return s.Name
}
