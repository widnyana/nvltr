package core

import (
	"io/ioutil"

	"github.com/widnyana/nvltr/conf"

	yaml "gopkg.in/yaml.v2"
)

// LoadConf provide load yml config.
func loadConf(confPath string) error {
	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(configFile), &conf.Config)

	if err != nil {
		return err
	}

	return nil
}
