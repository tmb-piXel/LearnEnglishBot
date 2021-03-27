package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/telegram"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	//	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	botAPI, err := tgbotapi.NewBotAPI(`1653360099:AAH7Xk3AU0HJYsb-B4sYX2qmu3MghuEDSM0`)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	dictionary := storage.ReadDictionary(cfg.DictionaryFile)
	bot := telegram.NewBot(botAPI, dictionary, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
