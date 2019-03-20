package filters

const (
	ValueSourceField = "field"
	ValueSourceConst = "const"
	ValueSourceAL    = "al"
	ValueSourceDict  = "dict"
)

type Config struct {
	Name     string          `yaml:"name,omitempty"`
	Not      bool            `yaml:"not,omitempty"`
	Sections []SectionConfig `yaml:"sections,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
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
	Field       string      `yaml:"field,omitempty"`
	Op          string      `yaml:"op,omitempty"`
	Value       interface{} `yaml:"value,omitempty"`
	ValueSource string      `yaml:"valueSource,omitempty"`
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
