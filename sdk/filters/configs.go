package filters

const (
	RVKindField = "field"
	RVKindConst = "const"
	RVKindDict  = "dict"
	RVKindAL    = "al"
)

type Config struct {
	Name     string          `yaml:"name,omitempty"`
	Not      bool            `yaml:"not,omitempty"`
	Join     bool            `yaml:"join,omitempty"`
	Sections []SectionConfig `yaml:"sections,omitempty"`
}

type SectionConfig struct {
	Or         bool              `yaml:"or,omitempty"`
	Not        bool              `yaml:"not,omitempty"`
	Conditions []ConditionConfig `yaml:"conditions,omitempty"`
}

type ConditionConfig struct {
	Lv     string `yaml:"lv,omitempty"`
	Op     string `yaml:"operator,omitempty"`
	Rv     string `yaml:"rv,omitempty"`
	RvKind string `yaml:"rv_kind,omitempty"`
}
