package dictionaries

import (
	"fmt"
)

type Dict map[string]string

type Config struct {
	Name string `json:"name"`
	Data Dict   `json:"data"`
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("dictionary: name required")
	}

	if len(r.Data) == 0 {
		return fmt.Errorf("dictionary: data required")
	}

	return nil
}
