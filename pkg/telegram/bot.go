package telegram

import (
	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	bot          *tb.Bot
	dictionaries map[string]map[string]string
	messages     config.Messages
}

func NewBot(bot *tb.Bot, dictionaries map[string]map[string]string, messages config.Messages) *Bot {
	return &Bot{
		bot:          bot,
		dictionaries: dictionaries,
		messages:     messages,
	}
}

func (b *Bot) Start() error {
	b.Handle()
	b.bot.Start()
	return nil
}
