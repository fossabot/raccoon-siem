package filters

const (
	RvSourceField = "field"
	RvSourceConst = "const"
	RvSourceAL    = "al"
	RvSourceDict  = "dict"
)

type Config struct {
	Name         string              `yaml:"name,omitempty"`
	Not          bool                `yaml:"not,omitempty"`
	Join         bool                `yaml:"join,omitempty"`
	Sections     []SectionConfig     `yaml:"sections,omitempty"`
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
	LeftEventID     string
	LeftEventField  string
	Operator        string
	RightEventID    string
	RightEventField string
}
