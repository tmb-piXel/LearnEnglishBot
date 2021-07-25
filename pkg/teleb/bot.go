package teleb

import (
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	bot        *tb.Bot
	dictionary map[string]string
	messages   config.Messages
}

func NewBot(bot *tb.Bot, dictionary map[string]string, messages config.Messages) *Bot {
	return &Bot{
		bot:        bot,
		dictionary: dictionary,
		messages:   messages,
	}
}

func (b *Bot) Start() error {
	b.Handle()
	b.bot.Start()
	return nil
}
