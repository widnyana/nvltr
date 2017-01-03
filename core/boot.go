package core

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/widnyana/nvltr/conf"
)

var (
	configLoc      = flag.String("c", "config.yml", "Path To Config File")
	configLocation string
	cfx            context.Context
)

// Boot create new nvltr instance
func Boot() error {
	flag.Parse()

	configLocation = string(*configLoc)

	abspath := filepath.IsAbs(*configLoc)
	if !abspath {
		cwd, _ := os.Getwd()
		configLocation = fmt.Sprintf("%s/%s", cwd, configLocation)
	}

	err := loadConf(configLocation)
	if err != nil {
		err = fmt.Errorf("Error Loading Config: %s", err.Error())
		return err
	}

	err = initLog()
	if err != nil {
		err = fmt.Errorf("Error Initting Log: %s", err.Error())
		return err
	}

	cfx = newCtxWithValues(conf.Config)
	return nil
}
