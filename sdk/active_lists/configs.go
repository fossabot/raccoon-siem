package activeLists

import "time"

const (
	OpGet        = "get"
	OpSet        = "set"
	OpDel        = "del"
	alNamePrefix = "raccoon-al-"
)

type Config struct {
	Name string        `yaml:"name"`
	TTL  time.Duration `yaml:"ttl"`
}

type Mapping struct {
	EventField string      `yaml:"eventField"`
	ALField    string      `yaml:"activeListField"`
	Constant   interface{} `yaml:"constant"`
}
