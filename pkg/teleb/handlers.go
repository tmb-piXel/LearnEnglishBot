package teleb

import (
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

var word = make(map[int64]string)
var IDofUserChats = make(map[int64]bool) //Map key - chatID, value - isEnteredStart

func (b *Bot) Handle() {

	var (
		menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		selector = &tb.ReplyMarkup{}
		st       = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")
		btnStart    = st.Text("Start")

		btnDE = selector.Data("🇩🇪 DE", "German")
		btnEN = selector.Data("🇬🇧 EN", "English")
	)

	st.Reply(
		st.Row(btnStart),
	)

	menu.Reply(
		menu.Row(btnHelp, btnSettings),
	)
	selector.Inline(
		selector.Row(btnDE, btnEN),
	)

	// enWord := make(map[int64]string)      //Map key - chatID, value - enWord

	// chatID := update.Message.Chat.ID

	// if !Contains(IDofUserChats, chatID) {
	// 	//Check if a new user
	// 	IDofUserChats[chatID] = false
	// 	b.startChat(chatID)
	// } else if update.Message.IsCommand() {
	// 	//Check the message is this command or not and processing
	// 	enWord[chatID] = GetRandomKey(b.dictionary)
	// 	isEnteredStart, _ = b.handleCommand(update.Message, enWord[chatID])
	// 	IDofUserChats[chatID] = isEnteredStart
	// 	if !IDofUserChats[chatID] {
	// 		b.startChat(chatID)
	// 		continue
	// 	}
	// } else {
	// 	//Processing messages
	// 	fmt.Println(IDofUserChats)
	// 	if IDofUserChats[chatID] {
	// 		b.checkAnswer(update.Message, enWord[chatID], b.dictionary)
	// 		enWord[chatID] = GetRandomKey(b.dictionary)
	// 		b.sendEnWord(update.Message, enWord[chatID])
	// 	} else {
	// 		b.startChat(chatID)
	// 		continue
	// 	}
	// }

	// Command: /start <PAYLOAD>

	// On reply button pressed (message)
	b.bot.Handle(&btnStart, func(m *tb.Message) {
		if !m.Private() {
			return
		}

		IDofUserChats[m.Chat.ID] = true
		b.bot.Send(m.Chat, "Добрый день! Выберите язык", selector)
	})

	b.bot.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		//IDofUserChats[m.Chat.ID] = true
		b.bot.Send(m.Chat, "Добрый день! Выберите язык", selector)
	})

	// On reply button pressed (message)
	b.bot.Handle(&btnHelp, func(m *tb.Message) {
		b.bot.Send(m.Chat, "helpmessage")
	})

	b.bot.Handle(&btnSettings, func(m *tb.Message) {
		b.bot.Send(m.Chat, "settingmessage", selector)
	})

	// On inline button pressed (callback)
	b.bot.Handle(&btnDE, func(c *tb.Callback) {

		b.dictionary = storage.ReadDictionary("dictionaries/german")
		word[c.Message.Chat.ID] = GetRandomKey(b.dictionary)
		b.bot.Send(c.Message.Chat, word[c.Message.Chat.ID], menu)

		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen German",
		})
	})

	b.bot.Handle(&btnEN, func(c *tb.Callback) {

		b.dictionary = storage.ReadDictionary("dictionaries/english")

		word[c.Message.Chat.ID] = GetRandomKey(b.dictionary)
		b.bot.Send(c.Message.Chat, word[c.Message.Chat.ID], menu)

		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen English",
		})
	})

	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		if IDofUserChats[m.Chat.ID] {
			b.bot.Send(m.Chat, word[m.Chat.ID])
			word[m.Chat.ID] = GetRandomKey(b.dictionary)
			b.bot.Send(m.Chat, word[m.Chat.ID])
		} else {
			b.bot.Send(m.Chat, "Нажмите старт", st)
		}

	})
}

// func (b *Bot) checkAnswer(message *tgbotapi.Message, enWord string, dictionary map[string]string) error {
// 	if Compaire(dictionary[enWord], message.Text) {
// 		err := b.sendMessage(message.Chat.ID, b.messages.CorrectAnswer)
// 		return err
// 	} else {
// 		err := b.sendMessage(message.Chat.ID, b.messages.WrongAnswer)
// 		if err != nil {
// 			return err
// 		}
// 		err = b.sendMessage(message.Chat.ID, b.messages.TheCorrectAnswerWas+dictionary[enWord])
// 		return err
// 	}
// }

// func (b *Bot) sendMessage(chatID int64, msg string) error {
// 	message := tgbotapi.NewMessage(chatID, msg)
// 	_, err := b.bot.Send(message)
// 	return err
// }
