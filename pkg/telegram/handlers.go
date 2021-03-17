package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) (isEnteredStart bool, err error) {
	isEnteredStart = false
	switch message.Command() {
	case commandStart:
		isEnteredStart = true
		err = b.handleStartCommand(message)
	default:
		err = b.handleUnknownCommand(message)
	}
	return isEnteredStart, err
}

func (b *Bot) startChat(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Введите /start")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "enWord")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "UnknownCommand")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "enWord")
	_, err := b.bot.Send(msg)
	return err
}
