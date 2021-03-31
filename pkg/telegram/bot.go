package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	dictionary map[string]string
	messages   config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, dictionary map[string]string, messages config.Messages) *Bot {
	return &Bot{
		bot:        bot,
		dictionary: dictionary,
		messages:   messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 600

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	var isEnteredStart bool
	IDofUserChats := make(map[int64]bool) //Map key - chatID, value - isEnteredStart
	enWord := make(map[int64]string)      //Map key - chatID, value - enWord

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if !Contains(IDofUserChats, chatID) {
			//Check if a new user
			IDofUserChats[chatID] = false
			b.startChat(chatID)
		} else if update.Message.IsCommand() {
			//Check the message is this command or not and processing commands
			enWord[chatID] = GetRandomKey(b.dictionary)
			isEnteredStart, _ = b.handleCommand(update.Message, enWord[chatID])
			IDofUserChats[chatID] = isEnteredStart
			if !IDofUserChats[chatID] {
				b.startChat(chatID)
				continue
			}
		} else {
			//Processing messages
			fmt.Println(IDofUserChats)
			if IDofUserChats[chatID] {
				b.checkAnswer(update.Message, enWord[chatID], b.dictionary)
				enWord[chatID] = GetRandomKey(b.dictionary)
				b.sendEnWord(update.Message, enWord[chatID])
			} else {
				b.startChat(chatID)
				continue
			}
		}
	}

	return nil
}
