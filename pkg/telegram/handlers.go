package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	setLanguage  = "set"
)

func (b *Bot) handleCommand(message *tgbotapi.Message, enWord string) (isEnteredStart bool, err error) {
	isEnteredStart = false
	switch message.Command() {
	case commandStart:
		isEnteredStart = true
		err = b.handleStartCommand(message, enWord)
	case setLanguage:
		b.handleSetLanguage(message)
	default:
		err = b.handleUnknownCommand(message)
	}
	return isEnteredStart, err
}

func (b *Bot) handleSetLanguage(message *tgbotapi.Message) {
	// c := message.CommandArguments()
	// b.dictionary = storage.ReadDictionary("dictionaries/" + c)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Select your class")

	btn := tgbotapi.NewInlineKeyboardButtonData("ste", "set")
	var row []tgbotapi.InlineKeyboardButton
	row = append(row, btn)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) startChat(chatID int64) error {
	err := b.sendMessage(chatID, b.messages.Start)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message, enWord string) error {
	err := b.sendMessage(message.Chat.ID, b.messages.AlreadyStart)
	if err != nil {
		return err
	}
	err = b.sendMessage(message.Chat.ID, enWord)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	err := b.sendMessage(message.Chat.ID, b.messages.UnknownCommand)
	return err
}

func (b *Bot) sendEnWord(message *tgbotapi.Message, enWord string) error {
	err := b.sendMessage(message.Chat.ID, enWord)
	return err
}

func (b *Bot) checkAnswer(message *tgbotapi.Message, enWord string, dictionary map[string]string) error {
	if Compaire(dictionary[enWord], message.Text) {
		err := b.sendMessage(message.Chat.ID, b.messages.CorrectAnswer)
		return err
	} else {
		err := b.sendMessage(message.Chat.ID, b.messages.WrongAnswer)
		if err != nil {
			return err
		}
		err = b.sendMessage(message.Chat.ID, b.messages.TheCorrectAnswerWas+dictionary[enWord])
		return err
	}
}

func (b *Bot) sendMessage(chatID int64, msg string) error {
	message := tgbotapi.NewMessage(chatID, msg)
	_, err := b.bot.Send(message)
	return err
}
