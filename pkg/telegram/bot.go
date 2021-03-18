package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/config"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{
		bot:      bot,
		messages: messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	var isEnteredStart bool
	IDofUserChats := make(map[int64]bool) //Map key - chatID, value - isEnteredStart
	enWord := make(map[int64]string)      //Map key - chatID, value - enWord
	dictionary := storage.ReadDictionary()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if contains(IDofUserChats, chatID) == false {
			//Check if a new user
			IDofUserChats[chatID] = false
			b.startChat(chatID)
		} else if update.Message.IsCommand() {
			//Check the message is this command or not and processing commands
			enWord[chatID] = getRandomKey(dictionary)
			isEnteredStart, _ = b.handleCommand(update.Message, enWord[chatID])
			IDofUserChats[chatID] = isEnteredStart
			fmt.Println(IDofUserChats)
			if IDofUserChats[chatID] != true {
				b.startChat(chatID)
				continue
			}
		} else {
			//Processing messages
			fmt.Println(IDofUserChats)
			if IDofUserChats[chatID] == true {
				b.checkAnswer(update.Message, enWord[chatID], dictionary)
				enWord[chatID] = getRandomKey(dictionary)
				b.sendEnWord(update.Message, enWord[chatID])
			} else {
				b.startChat(chatID)
				continue
			}
		}
	}

	return nil
}
