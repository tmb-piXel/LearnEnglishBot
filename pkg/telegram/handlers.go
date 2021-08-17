package telegram

import (
	"fmt"
	"strings"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

//TODO make dictionariesInterface
//TODO refactor handlers.go
//TODO create model user and dboUsers

var IsTheStartPressed = make(map[int64]bool) //Is the start pressed for chatID

type user struct {
	chatID        int64
	isForeignToRU bool
	original      *[]string
	translated    *[]string
	iWord         int
}

var users = make(map[int64]user)

func (b *Bot) Handle() {
	var (
		menuMarkup  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		startMarkup = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		langMarkup  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		modeMarkup  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

		settingsBtn  = menuMarkup.Text("⚙ Settings")
		startBtn     = startMarkup.Text("Start")
		listBtn      = modeMarkup.Data("List", "List")
		ruToBtn      = modeMarkup.Data("ruTo", "tuRu")
		toRuBtn      = modeMarkup.Data("toRu", "toRu")
		topicsMarkup = make(map[string]*tb.ReplyMarkup)
		langBtns     []tb.Btn
		topicBtns    = make(map[string][]tb.Btn)
	)

	languages := storage.GetLanguages()

	//Set buttons and mapkup
	for _, l := range languages {
		codeLanguages := storage.GetCode(l)
		langBtns = append(langBtns, langMarkup.Data(codeLanguages, l))
		topicTitles := storage.GetTopicTitles(l)
		topicsMarkup[l] = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		for _, t := range topicTitles {
			topicBtns[l] = append(topicBtns[l], topicsMarkup[l].Data(t, l+"_"+t))
		}
		topicsMarkup[l].Inline(topicsMarkup[l].Row(topicBtns[l]...))
	}

	startMarkup.Reply(startMarkup.Row(startBtn))
	menuMarkup.Reply(menuMarkup.Row(settingsBtn))

	langMarkup.Inline(
		langMarkup.Row(langBtns[0:len(langBtns)/2]...),
		langMarkup.Row(langBtns[len(langBtns)/2:]...),
	)

	modeMarkup.Inline(
		langMarkup.Row(listBtn, ruToBtn, toRuBtn),
	)

	//Handle start button
	b.bot.Handle(&startBtn, func(m *tb.Message) {
		if !m.Private() {
			return
		}
		IsTheStartPressed[m.Chat.ID] = true
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)
	})

	//Handel setting buttons
	b.bot.Handle(&settingsBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)
	})

	//Buttons selected language
	for _, button := range langBtns {
		btn := button
		callback := func(c *tb.Callback) {
			b.bot.Respond(c, &tb.CallbackResponse{
				Text: "You have chosen " + btn.Unique,
			})
			b.bot.Send(c.Message.Chat, "Выберите тему", topicsMarkup[btn.Unique])
		}
		b.bot.Handle(&btn, callback)
	}

	//Buttons selected topic
	for _, buttons := range topicBtns {
		for _, button := range buttons {
			btn := button
			lang := strings.Split(btn.Unique, "_")[0]
			callback := func(c *tb.Callback) {
				b.bot.Respond(c, &tb.CallbackResponse{
					Text: "You have chosen " + btn.Text,
				})
				user := users[c.Message.Chat.ID]
				user.original = storage.GetOriginalWords(lang, btn.Text)
				user.translated = storage.GetTransletedWords(lang, btn.Text)
				users[c.Message.Chat.ID] = user
				b.bot.Send(c.Message.Chat, "Выберите режим", modeMarkup)
			}
			b.bot.Handle(&btn, callback)
		}
	}

	//Handle List
	b.bot.Handle(&listBtn, func(c *tb.Callback) {
		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen " + listBtn.Unique,
		})

		chatID := c.Message.Chat.ID
		var list []string
		for i, w := range *users[chatID].original {
			list = append(list, w+"-"+(*users[chatID].translated)[i])
		}
		_, err := b.bot.Send(c.Message.Chat, strings.Join(list, "\n"), menuMarkup)
		if err != nil {
			fmt.Println(err)
		}
	})

	//Handle ruTo
	b.bot.Handle(&ruToBtn, func(c *tb.Callback) {
		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen " + ruToBtn.Unique,
		})

		chatID := c.Message.Chat.ID
		user := users[chatID]
		user.isForeignToRU = false
		user.iWord = GetR(*user.translated)
		word := (*user.translated)[user.iWord]
		b.bot.Send(c.Message.Chat, word, menuMarkup)
		users[chatID] = user
	})

	//Handle toRu
	b.bot.Handle(&toRuBtn, func(c *tb.Callback) {
		b.bot.Respond(c, &tb.CallbackResponse{
			Text: "You have chosen " + toRuBtn.Unique,
		})

		chatID := c.Message.Chat.ID
		user := users[chatID]
		user.isForeignToRU = true
		user.iWord = GetR(*user.original)
		word := (*user.original)[user.iWord]
		b.bot.Send(c.Message.Chat, word, menuMarkup)
		users[chatID] = user
	})

	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		chatID := m.Chat.ID
		if IsTheStartPressed[chatID] {
			user := users[chatID]
			i := user.iWord

			var first []string
			var second []string
			if user.isForeignToRU {
				first = *user.original
				second = *user.translated
			} else {
				first = *user.translated
				second = *user.original
			}

			if CheckAnswer(second[i], m.Text) {
				b.bot.Send(m.Chat, b.messages.CorrectAnswer)
			} else {
				b.bot.Send(m.Chat, b.messages.WrongAnswer)
				b.bot.Send(m.Chat, b.messages.TheCorrectAnswerWas+second[i])
			}

			i = GetR(first)
			b.bot.Send(m.Chat, first[i])
			user.iWord = i
			users[m.Chat.ID] = user
		} else {
			users[chatID] = user{chatID: chatID}
			b.bot.Send(m.Chat, b.messages.StartMessage, startMarkup)
		}
	})
}
