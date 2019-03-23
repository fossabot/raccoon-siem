package enrichment

const (
	ValueSourceKindConst = "constant"
	ValueSourceKindDict  = "dictionary"
	ValueSourceKindEvent = "event"
	ValueSourceKindAL    = "activeList"
)

type Config struct {
	Field            string      `yaml:"field,omitempty"`
	Constant         interface{} `yaml:"constant,omitempty"`
	KeyFields        []string    `yaml:"keyFields,omitempty"`
	ValueSourceKind  string      `yaml:"valueSourceKind,omitempty"`
	ValueSourceName  string      `yaml:"valueSourceName,omitempty"`
	ValueSourceField string      `yaml:"valueSourceField,omitempty"`
	TriggerField     string      `yaml:"trigger_field,omitempty"`
	TriggerValue     interface{} `yaml:"trigger_value,omitempty"`
}
