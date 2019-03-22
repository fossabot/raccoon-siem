package activeLists

import "time"

const (
	OpGet        = "get"
	OpSet        = "set"
	OpDel        = "del"
	alNamePrefix = "al-"
)

type Config struct {
	Name string        `yaml:"name"`
	TTL  time.Duration `yaml:"ttl"`
}

type Mapping struct {
	EventField string
	ALField    string
}
