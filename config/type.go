package config

type Config struct {
	//	TODO:
	Gitlab struct {
		Token         string `yaml:"token"`
		WebhookHost   string `yaml:"webhook_host"`
		WebhookMethod string `yaml:"webhook_method"`
		WebhookPort   int    `yaml:"webhook_port"`
	} `yaml:"gitlab"`
}

// TODO: валидация
