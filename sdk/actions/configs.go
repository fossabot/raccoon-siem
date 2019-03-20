package actions

const (
	KindRelease    = "release"
	KindActiveList = "al"
)

const (
	ValueSourceKindConst = "const"
	ValueSourceKindEvent = "event"
	ValueSourceKindAL    = "al"
)

type ReleaseConfig struct {
	MutateConfigs []MutateConfig `yaml:"mutate,omitempty"`
}

type MutateConfig struct {
	Field            string   `yaml:"targetField,omitempty"`
	KeyFields        []string `yaml:"keyFields,omitempty"`
	Constant         string   `yaml:"constant,omitempty"`
	ValueSourceKind  string   `yaml:"valueSourceKind,omitempty"`
	ValueSourceName  string   `yaml:"valueSourceName,omitempty"`
	ValueSourceField string   `yaml:"valueSourceField,omitempty"`
}
