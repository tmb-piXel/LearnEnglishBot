package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/telegram"
)

const telegramToken = `1653360099:AAEidSka74r1KJtq9nzgpoZFEfeZbnfeyvQ`

func main() {
	botAPI, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot := telegram.NewBot(botAPI)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
