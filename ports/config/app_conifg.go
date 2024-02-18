package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config AppConfig

type AppConfig struct {
	MailSender MailSender
}

type MailSender struct {
	From               string `mapstructure:"from"`
	SMTP               string `mapstructure:"smtp"`
	Port               int    `mapstructure:"port"`
	Password           string `mapstructure:"password"`
	InsecureSkipVerify bool   `mapstructure:"insecure-skip-verify"`
}

func InitAppConfig() {
	var mailSend MailSender
	if err := viper.UnmarshalKey("mail-sender", &mailSend); err != nil {
		panic(fmt.Errorf("failed to load up mail sender config: %v", err))
	}

	Config.MailSender = mailSend
}
