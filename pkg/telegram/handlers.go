package telegram

import (
	"fmt"

	s "github.com/tmb-piXel/LearnEnglishBot/pkg/services"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

//TODO make dictionariesInterface
//TODO refactor handlers.go
//TODO Ограничить список слов 400 симвалами
//TODO Наполнить словари
//TODO Юнит тесты
//TODO Добавление БД
//TODO Логирование

func (b *Bot) Handle() {
	var (
		menuMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		langMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		modeMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		topicsMarkup = make(map[string]*tb.ReplyMarkup)

		settingsBtn = menuMarkup.Text("⚙ Настройки")
		helpBtn     = menuMarkup.Text("Помощь")
		setLangBtn  = modeMarkup.Text("Выбрать язык")
		setTopicBtn = modeMarkup.Text("Выбрать тему")
		listBtn     = modeMarkup.Text("Список слов")
		fromRuBtn   = modeMarkup.Text("Перевод с русского")
		toRuBtn     = modeMarkup.Text("Перевод на русский")
		topicBtns   = make(map[string][]tb.Btn)
		langBtns    []tb.Btn
	)

	languages := storage.GetLanguages()

	//Set lang buttons and topics markup
	for _, l := range languages {
		lang := l[8:] // delete flag
		langBtn := langMarkup.Data(l, lang)
		langBtns = append(langBtns, langBtn)

		topicTitles := storage.GetTopicTitles(l)
		topicsMarkup[lang] = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		for _, t := range topicTitles {
			topicBtn := topicsMarkup[lang].Data(t, lang+t)
			topicBtns[lang] = append(topicBtns[lang], topicBtn)
		}
		topicsMarkup[lang].Inline(topicsMarkup[lang].Split(1, topicBtns[lang])...)
	}

	menuMarkup.Reply(menuMarkup.Row(settingsBtn, helpBtn))
	modeMarkup.Reply(
		modeMarkup.Row(listBtn),
		modeMarkup.Row(setLangBtn, setTopicBtn),
		modeMarkup.Row(fromRuBtn, toRuBtn),
	)
	langMarkup.Inline(langMarkup.Split(1, langBtns)...)

	//Handle start button
	b.bot.Handle("/start", func(m *tb.Message) {
		s.NewUser(m.Chat.ID)
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)
	})

	//Handel setting button
	b.bot.Handle(&settingsBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Настройки", modeMarkup)
	})

	//Handel help button
	b.bot.Handle(&helpBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Помощь")
	})

	//Buttons selected language
	for _, button := range langBtns {
		btn := button
		callback := func(c *tb.Callback) {
			b.bot.Respond(c, &tb.CallbackResponse{
				Text: "You have chosen " + btn.Unique,
			})
			s.SetLanguage(c.Message.Chat.ID, btn.Text)
			b.bot.Send(c.Message.Chat, "Выберите тему", topicsMarkup[btn.Unique])
		}
		b.bot.Handle(&btn, callback)
	}

	//Buttons selected topic
	for _, buttons := range topicBtns {
		for _, button := range buttons {
			btn := button
			callback := func(c *tb.Callback) {
				b.bot.Respond(c, &tb.CallbackResponse{
					Text: "You have chosen " + btn.Text,
				})
				s.SetTopic(c.Message.Chat.ID, btn.Text)
				b.bot.Send(c.Message.Chat, s.NewWord(c.Message.Chat.ID), menuMarkup)
			}
			b.bot.Handle(&btn, callback)
		}
	}

	//Handle List
	b.bot.Handle(&listBtn, func(m *tb.Message) {
		_, err := b.bot.Send(m.Chat, s.ListWords(m.Chat.ID), menuMarkup)
		if err != nil {
			fmt.Println(err)
		}
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)
	})

	//Handel setting language buttons
	b.bot.Handle(&setLangBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)
	})

	//Handel setting topics buttons
	b.bot.Handle(&setTopicBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Выберите тему", topicsMarkup[s.Language(m.Chat.ID)[8:]])
	})

	//Handle ruTo
	b.bot.Handle(&fromRuBtn, func(m *tb.Message) {
		s.SetIsToRu(m.Chat.ID, false)
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)
	})

	//Handle toRu
	b.bot.Handle(&toRuBtn, func(m *tb.Message) {
		s.SetIsToRu(m.Chat.ID, true)
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)
	})

	//Handle text message
	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		chatID := m.Chat.ID
		word := s.Word(chatID)
		if CheckAnswer(word, m.Text) {
			b.bot.Send(m.Chat, b.messages.CorrectAnswer)
		} else {
			b.bot.Send(m.Chat, b.messages.WrongAnswer)
			b.bot.Send(m.Chat, b.messages.TheCorrectAnswerWas+word)
		}

		b.bot.Send(m.Chat, s.NewWord(chatID))
	})
}
