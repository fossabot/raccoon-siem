package sdk

type DefaultComponentSettings struct {
	Name        string `yaml:"name,omitempty"`
	Debug       bool   `yaml:"debug,omitempty"`
	MeterPeriod int    `yaml:"meterPeriod,omitempty"`
}
