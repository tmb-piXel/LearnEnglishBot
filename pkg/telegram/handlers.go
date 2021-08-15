package telegram

import (
	"strings"

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

		btnSettings = menu.Text("⚙ Settings")
		btnStart    = startMarkup.Text("Start")
		buttons     []tb.Btn
	)

	for title := range b.dictionaries {
		codeName := strings.Split(title, "_")
		buttons = append(buttons, setlang.Data(codeName[0], codeName[1]))
	}

	startMarkup.Reply(
		startMarkup.Row(btnStart),
	)
	menu.Reply(
		menu.Row(btnSettings),
	)
	setlang.Inline(
		setlang.Row(buttons...),
	)

	//Buttons selected language
	for _, button := range buttons {
		btn := button
		callback := func(c *tb.Callback) {
			chatID := c.Message.Chat.ID
			dictionary[chatID] = b.dictionaries[btn.Text+"_"+btn.Unique]
			words[chatID] = GetRandomWord(dictionary[chatID])
			b.bot.Send(c.Message.Chat, words[chatID], menu)

			b.bot.Respond(c, &tb.CallbackResponse{
				Text: "You have chosen " + btn.Unique,
			})
		}
		b.bot.Handle(&btn, callback)
	}

	b.bot.Handle(&btnStart, func(m *tb.Message) {
		if !m.Private() {
			return
		}
		IsTheStartPressed[m.Chat.ID] = true
		b.bot.Send(m.Chat, b.messages.SelectLanguage, setlang)
	})

	b.bot.Handle(&btnSettings, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.SelectLanguage, setlang)
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
