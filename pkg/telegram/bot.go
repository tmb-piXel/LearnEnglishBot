package telegram

import (
	"github.com/tmb-piXel/LearnEnglishBot/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	bot      *tb.Bot
	messages config.Messages
	buttons  config.Buttons
}

func NewBot(bot *tb.Bot, messages config.Messages, buttons config.Buttons) *Bot {
	return &Bot{
		bot:      bot,
		messages: messages,
		buttons:  buttons,
	}
}

func (b *Bot) Start() error {
	b.Handle()
	b.bot.Start()
	return nil
}
