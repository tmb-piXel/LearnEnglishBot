package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/telegram"
)

//const telegramToken = `1653360099:AAEidSka74r1KJtq9nzgpoZFEfeZbnfeyvQ`

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot := telegram.NewBot(botAPI, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
