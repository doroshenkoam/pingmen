package config

type Config struct {
	Gitlab struct {
		Token         string `yaml:"token"`
		WebhookHost   string `yaml:"webhook_host"`
		WebhookMethod string `yaml:"webhook_method"`
		WebhookPort   int    `yaml:"webhook_port"`

		Actions []string `yaml:"actions"`
	} `yaml:"gitlab"`
	Telegram struct {
		Token string `yaml:"token"`
		Debug bool   `yaml:"debug"`
	} `yaml:"telegram"`
	Users struct {
		Dictionary []string `yaml:"dictionary"`
	} `yaml:"users"`
	Projects struct {
		Dictionary []string `yaml:"dictionary"`
	} `yaml:"projects"`
}

// TODO: валидация
