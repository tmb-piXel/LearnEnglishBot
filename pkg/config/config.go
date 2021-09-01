package config

import (
	"github.com/spf13/viper"
)

type Messages struct {
	Responses
}

type Responses struct {
	StartMessage        string `mapstructure:"start_message"`
	AlreadyStart        string `mapstructure:"already_start"`
	UnknownCommand      string `mapstructure:"unknown_command"`
	CorrectAnswer       string `mapstructure:"correct_answer"`
	WrongAnswer         string `mapstructure:"wrong_answer"`
	TheCorrectAnswerWas string `mapstructure:"the_correct_answer_was"`
	SelectLanguage      string `mapstructure:"select_language"`
}

type Config struct {
	TelegramToken   string
	PathDictonaries string `mapstructure:"path_dictionaries"`
	PostgresqlUrl   string
	Messages        Messages
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
	if err := viper.BindEnv("telegramtoken", "postgresqlurl"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegramtoken")
	cfg.TelegramToken = `1653360099:AAHu1tOgMly0DFA-KecC7CeWuhnaF-f9_j8`
	cfg.PostgresqlUrl = viper.GetString("postgresqlurl")
	cfg.PostgresqlUrl = `host=localhost port=5432 user=pixel password=passwordpolyglot dbname=polyglot sslmode=disable`

	return nil
}
