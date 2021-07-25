package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/teleb"
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

	dictionary := storage.ReadDictionary(cfg.DictionaryFile)
	bot := teleb.NewBot(botAPI, dictionary, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
