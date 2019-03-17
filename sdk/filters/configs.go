package filters

const (
	RvSourceField = "field"
	RvSourceConst = "const"
	RvSourceAL    = "al"
	RvSourceDict  = "dict"
)

type Config struct {
	Name     string          `yaml:"name,omitempty"`
	Not      bool            `yaml:"not,omitempty"`
	Sections []SectionConfig `yaml:"sections,omitempty"`
}

type JoinConfig struct {
	Name         string              `yaml:"name,omitempty"`
	Not          bool                `yaml:"not,omitempty"`
	JoinSections []JoinSectionConfig `yaml:"joinSections,omitempty"`
}

type SectionConfig struct {
	Or         bool              `yaml:"or,omitempty"`
	Not        bool              `yaml:"not,omitempty"`
	Conditions []ConditionConfig `yaml:"conditions,omitempty"`
}

type ConditionConfig struct {
	Field    string      `yaml:"field,omitempty"`
	Op       string      `yaml:"op,omitempty"`
	Rv       interface{} `yaml:"rv,omitempty"`
	RvSource string      `yaml:"rvKind,omitempty"`
}

type JoinSectionConfig struct {
	Or         bool                  `yaml:"or,omitempty"`
	Not        bool                  `yaml:"not,omitempty"`
	Conditions []JoinConditionConfig `yaml:"conditions,omitempty"`
}

type JoinConditionConfig struct {
	LeftTag    string `yaml:"leftTag,omitempty"`
	LeftField  string `yaml:"leftField,omitempty"`
	Op         string `yaml:"op,omitempty"`
	RightTag   string `yaml:"rightTag,omitempty"`
	RightField string `yaml:"rightField,omitempty"`
}
