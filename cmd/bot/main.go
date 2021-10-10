package main

import (
	"time"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/db"
	log "github.com/tmb-piXel/LearnEnglishBot/pkg/logger"
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

	if err != nil {
		log.Panic("Bot did not start error: ", err)
	}

	log.Println("Bot started")

	db.InitDB(cfg.PostgresqlUrl)
	storage.InitDictionaries(cfg.PathDictonaries)
	bot := telegram.NewBot(botAPI, cfg.Messages, cfg.Buttons)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
