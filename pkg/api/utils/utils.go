package utils

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// ReadConfig which will read a configuration yml and marshall it
// into an object.
func ReadConfig(configFile string) (*ServiceConfig, error) {
	file, err := ioutil.ReadFile(configFile) // just pass the file name

	if err != nil {
		return nil, err
	}

	var config ServiceConfig
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
