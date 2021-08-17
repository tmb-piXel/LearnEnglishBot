package telegram

import (
	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	bot      *tb.Bot
	messages config.Messages
}

func NewBot(bot *tb.Bot, messages config.Messages) *Bot {
	return &Bot{
		bot:      bot,
		messages: messages,
	}
}

func (b *Bot) Start() error {
	b.Handle()
	b.bot.Start()
	return nil
}
