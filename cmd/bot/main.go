package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tb.NewBot(tb.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 60 * time.Second},
	})

	log.Printf("Authorized on account %s", botAPI.URL)

	if err != nil {
		fmt.Printf("Bot did not start error: %s", err)
	}

	dictionaries := storage.ReadDictionaries(cfg.DictionaryFile)
	bot := telegram.NewBot(botAPI, dictionaries, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
