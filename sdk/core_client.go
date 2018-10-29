package sdk

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func CoreQuery(url string, dst interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(body, dst); err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%s", string(body))
		}
		return err
	}

	return nil
}
