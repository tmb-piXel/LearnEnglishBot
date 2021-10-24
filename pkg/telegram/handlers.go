package telegram

import (
	"fmt"
	"strconv"

	log "github.com/tmb-piXel/LearnEnglishBot/pkg/logger"
	s "github.com/tmb-piXel/LearnEnglishBot/pkg/services"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

//TODO make dictionariesInterface
//TODO refactor handlers.go
//TODO Наполнить словари
//TODO Юнит тесты
//TODO Логирование

func (b *Bot) Handle() {
	var (
		menuMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		langMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		modeMarkup   = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		topicsMarkup = make(map[string]*tb.ReplyMarkup)

		settingsBtn = menuMarkup.Text(b.buttons.Settings)
		helpBtn     = menuMarkup.Text(b.buttons.Help)
		setLangBtn  = modeMarkup.Text(b.buttons.SetLang)
		setTopicBtn = modeMarkup.Text(b.buttons.SetTopic)
		listBtn     = modeMarkup.Text(b.buttons.List)
		fromRuBtn   = modeMarkup.Text(b.buttons.FromRu)
		toRuBtn     = modeMarkup.Text(b.buttons.ToRu)
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

		var rows []tb.Row

		for i, t := range topicTitles {
			topicBtn := topicsMarkup[lang].Data(t, lang+fmt.Sprintf("%d", i))
			topicBtns[lang] = append(topicBtns[lang], topicBtn)

			row := langMarkup.Row(topicBtn)
			rows = append(rows, row)
		}

		topicsMarkup[lang].Inline(rows...)
	}

	menuMarkup.Reply(menuMarkup.Row(settingsBtn, helpBtn))
	modeMarkup.Reply(
		modeMarkup.Row(listBtn),
		modeMarkup.Row(setLangBtn, setTopicBtn),
		modeMarkup.Row(fromRuBtn, toRuBtn),
	)

	//List lang buttons
	var rows []tb.Row
	for _, b := range langBtns {
		row := langMarkup.Row(b)
		rows = append(rows, row)
	}
	langMarkup.Inline(rows...)

	//Handle start button
	b.bot.Handle("/start", func(m *tb.Message) {
		s.NewUser(m.Chat.ID)
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)

		log.Printf("Handle start",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handel setting button
	b.bot.Handle(&settingsBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Настройки", modeMarkup)

		log.Printf("Handle settings",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handel help button
	b.bot.Handle(&helpBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.HelpMessage)

		log.Printf("Handle Help",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
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
					Text: "You have chosen topic " + btn.Text,
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
			log.Error(err)
		}
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)

		log.Printf("Handle List",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handel setting language buttons
	b.bot.Handle(&setLangBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, b.messages.SelectLanguage, langMarkup)

		log.Printf("Handle setLanguage",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handel setting topics buttons
	b.bot.Handle(&setTopicBtn, func(m *tb.Message) {
		b.bot.Send(m.Chat, "Выберите тему", topicsMarkup[s.Language(m.Chat.ID)[8:]])

		log.Printf("Handle setTopic",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handle ruTo
	b.bot.Handle(&fromRuBtn, func(m *tb.Message) {
		s.SetIsToRu(m.Chat.ID, false)
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)

		log.Printf("Handle fromRu",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handle toRu
	b.bot.Handle(&toRuBtn, func(m *tb.Message) {
		s.SetIsToRu(m.Chat.ID, true)
		b.bot.Send(m.Chat, s.NewWord(m.Chat.ID), menuMarkup)

		log.Printf("Handle toRu",
			strconv.FormatInt(m.Chat.ID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
		)
	})

	//Handle text message
	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		chatID := m.Chat.ID
		word := s.Word(chatID)

		log.Printf("Transleted: %s, UserWord: %s",
			strconv.FormatInt(chatID, 10),
			m.Chat.FirstName+" "+m.Chat.LastName,
			word,
			m.Text,
		)

		if CheckAnswer(word, m.Text) {
			b.bot.Send(m.Chat, b.messages.CorrectAnswer+"\n\n"+s.NewWord(chatID))
		} else {
			b.bot.Send(m.Chat, b.messages.WrongAnswer+word+"\n\n"+s.NewWord(chatID))
		}
	})
}
