package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Load - loading and parse config
func Load(path string, cfg *Config) error {
	log.Printf("Config:Load: start")
	defer log.Printf("Config:Load: inited")

	log.Printf("Config:Load: Config file is: %s", path)

	cfgData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		return err
	}

	if err := cfg.validate(); err != nil {
		return err
	}

	return nil
}

// validate - configuration validator
func (cfg *Config) validate() error {
	if cfg.Telegram.Token == "" {
		return errForFields("telegram token")
	}

	if cfg.Telegram.ChatID == 0 {
		return errForFields("telegram chat_id")
	}

	if cfg.Telegram.WorkersCount == 0 {
		return errForFields("telegram workers_count")
	}

	if cfg.Gitlab.WebhookPort == 0 {
		return errForFields("gitlab webhook port")
	}

	if cfg.Gitlab.Token == "" {
		return errForFields("gitlab token")
	}

	if cfg.Gitlab.WebhookMethod == "" {
		return errForFields("gitlab webhook method")
	}

	if len(cfg.Gitlab.Actions) == 0 {
		return errForFields("gitlab actions")
	}

	if len(cfg.Users.Dictionary) == 0 {
		return errForFields("users dictionary")
	}

	if len(cfg.Projects.Dictionary) == 0 {
		return errForFields("projects dictionary")
	}

	return nil
}

// errForFields - return errors for fields
func errForFields(field string) error {
	return fmt.Errorf("error: configuration param %s is incorrect", field)
}
