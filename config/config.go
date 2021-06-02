package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Load - loading and parse config
func Load(path *string, cfg *Config) error {
	log.Printf("Config:Load: start")
	defer log.Printf("Config:Load: inited")

	cfgPath := *path
	log.Printf("Config:Load: Config file is: %s", cfgPath)

	cfgData, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		return err
	}

	return nil
}
