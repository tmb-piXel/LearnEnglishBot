package main

import (
	"log"
	"time"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/db"
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

	log.Printf("%s started", botAPI.Me.FirstName)

	if err != nil {
		log.Printf("Bot did not start error: %s", err)
	}

	db.InitDB(cfg.PostgresqlUrl)
	storage.InitDictionaries(cfg.PathDictonaries)
	bot := telegram.NewBot(botAPI, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
