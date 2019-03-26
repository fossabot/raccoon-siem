package activeLists

import "time"

const (
	OpGet        = "get"
	OpSet        = "set"
	OpDel        = "del"
	alNamePrefix = "raccoon-al-"
)

type Config struct {
	Name string `yaml:"name"`
	TTL  int64  `yaml:"ttl"`
}

func (r *Config) TTLDuration() time.Duration {
	return time.Duration(r.TTL) * time.Second
}

type Mapping struct {
	EventField string      `yaml:"eventField"`
	ALField    string      `yaml:"activeListField"`
	Constant   interface{} `yaml:"constant"`
}
