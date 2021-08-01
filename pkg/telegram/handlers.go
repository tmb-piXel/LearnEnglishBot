package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var dictionary = make(map[int64]map[string]string) //Dictionary for chatID
var words = make(map[int64]string)                 //Words for chatID
var IsTheStartPressed = make(map[int64]bool)       //Is the start pressed for chatID

func (b *Bot) Handle() {
	var (
		menu        = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		startMarkup = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		setlang     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")
		btnStart    = startMarkup.Text("Start")

		btnDE = setlang.Data("🇩🇪 DE", "German")
		btnEN = setlang.Data("🇬🇧 EN", "English")
	)

	startMarkup.Reply(
		startMarkup.Row(btnStart),
	)
	menu.Reply(
		menu.Row(btnHelp, btnSettings),
	)
	setlang.Inline(
		setlang.Row(btnDE, btnEN),
	)

	b.bot.Handle(&btnStart, func(m *tb.Message) {
		if !m.Private() {
			return
		}
		IsTheStartPressed[m.Chat.ID] = true
		b.bot.Send(m.Chat, b.messages.SelectLanguage, setlang)
	})

	b.bot.Handle(&btnHelp, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Help")
	})

	b.bot.Handle(&btnSettings, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.SelectLanguage, setlang)
	})

	b.bot.Handle(&btnDE, func(c *tb.Callback) {
		chatID := c.Message.Chat.ID
		dictionary[chatID] = b.dictionaries["german"]
		words[chatID] = GetRandomWord(dictionary[chatID])
		b.bot.Send(c.Message.Chat, words[chatID], menu)

		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen German",
		})
	})

	b.bot.Handle(&btnEN, func(c *tb.Callback) {
		chatID := c.Message.Chat.ID
		dictionary[chatID] = b.dictionaries["english"]
		words[chatID] = GetRandomWord(dictionary[chatID])
		b.bot.Send(c.Message.Chat, words[chatID], menu)

		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen English",
		})
	})

	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		if IsTheStartPressed[m.Chat.ID] {
			originalWord := dictionary[m.Chat.ID][words[m.Chat.ID]]
			if CheckAnswer(originalWord, m.Text) {
				b.bot.Send(m.Chat, b.messages.CorrectAnswer)
			} else {
				b.bot.Send(m.Chat, b.messages.WrongAnswer)
				b.bot.Send(m.Chat, b.messages.TheCorrectAnswerWas+originalWord)
			}
			words[m.Chat.ID] = GetRandomWord(dictionary[m.Chat.ID])
			b.bot.Send(m.Chat, words[m.Chat.ID])
		} else {
			b.bot.Send(m.Chat, b.messages.StartMessage, startMarkup)
		}
	})
}
