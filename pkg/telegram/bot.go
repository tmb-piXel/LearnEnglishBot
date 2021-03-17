package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	var IDofUserChats []int64
	var isEnteredStart bool

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if contains(IDofUserChats, update.Message.Chat.ID) == false {
			//Check if a new user
			IDofUserChats = append(IDofUserChats, update.Message.Chat.ID)
			b.startChat(update.Message.Chat.ID)
		} else if update.Message.IsCommand() {
			//Check the message is this command or not and processing commands
			if isEnteredStart, _ = b.handleCommand(update.Message); isEnteredStart != true {
				b.startChat(update.Message.Chat.ID)
				continue
			}
		} else {
			//Processing messages
			if isEnteredStart == true {
				b.handleMessage(update.Message)
			} else {
				b.startChat(update.Message.Chat.ID)
				continue
			}
		}
	}

	return nil
}
