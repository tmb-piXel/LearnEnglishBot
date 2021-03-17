package telegram

// This is dictionary for learning English Words

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
)

const token = `1653360099:AAEidSka74r1KJtq9nzgpoZFEfeZbnfeyvQ`

//StartBot - Run telegram bot
func StartBot() {
	bot := createBot()
	updates := getUpdates(bot)
	dictionary := storage.ReadDictionary()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		sendMessage(bot, chatID, "Введите /start")
		break
	}
	sendWords(bot, updates, dictionary)
}

// Send word from dictionary in chat
func sendWords(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, dictionary map[string]string) {
	var chatID int64
	for englishWord := range dictionary {
		// Messege with new word
		sendMessage(bot, chatID, englishWord)
		for update := range updates {
			chatID = update.Message.Chat.ID
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Check response
			switch update.Message.Text {
			case "/start":
				sendMessage(bot, chatID, englishWord)
				continue
			default:
				resp := update.Message.Text
				if strings.Contains(dictionary[englishWord], resp) {
					sendMessage(bot, chatID, "This is rigth answer")
				} else {
					sendMessage(bot, chatID, "Oooops, this is wrong answer")
					sendMessage(bot, chatID, "Right answer is: "+dictionary[englishWord])
				}
			}
			break
		}
	}
}

// Send message in chat
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, msg string) {
	newMsg := tgbotapi.NewMessage(chatID, msg)
	bot.Send(newMsg)
}

// Create telegram bot
func createBot() (bot *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

// Get updates bot
func getUpdates(bot *tgbotapi.BotAPI) (updates tgbotapi.UpdatesChannel) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	return updates
}
