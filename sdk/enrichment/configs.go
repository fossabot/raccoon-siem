package enrichment

const (
	FromConst = "const"
	FromDict  = "dict"
	FromAL    = "al"
)

type Config struct {
	Field            string   `yaml:"field,omitempty"`
	Constant         string   `yaml:"constant,omitempty"`
	KeyFields        []string `yaml:"keyFields,omitempty"`
	ValueSourceKind  string   `yaml:"valueSourceKind,omitempty"`
	ValueSourceName  string   `yaml:"valueSourceName,omitempty"`
	ValueSourceField string   `yaml:"valueSourceField,omitempty"`
	TriggerField     string   `yaml:"triggerField,omitempty"`
	TriggerValue     string   `yaml:"triggerValue,omitempty"`
}
