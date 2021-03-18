package config

import (
	"os"

	"github.com/spf13/viper"
)

type Messages struct {
	Responses
}

type Responses struct {
	Start               string `mapstructure:"start"`
	AlreadyStart        string `mapstructure:"already_start"`
	UnknownCommand      string `mapsturcture:"unknown_command"`
	CorrectAnswer       string `mapstructure:"correct_answer"`
	WrongAnswer         string `mapstructure:"wrong_answer"`
	TheCorrectAnswerWas string `mapsturcture:"theCorrectAnswerWas"`
}

type Config struct {
	TelegramToken string
	Messages      Messages
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TELEGRAMTOKEN", "1653360099:AAEidSka74r1KJtq9nzgpoZFEfeZbnfeyvQ")
	if err := viper.BindEnv("telegramtoken"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegramtoken")

	return nil
}
