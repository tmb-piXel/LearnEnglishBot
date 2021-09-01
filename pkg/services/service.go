package services

import (
	"math/rand"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/db"
	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
)

func NewUser(chatID int64) {
	u := m.NewUser(chatID)
	db.SaveUser(u)
}

func NewWord(chatID int64) (word string) {
	u := db.FindUser(chatID)
	words := storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	if u.GetIsToRu() {
		words = storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	}
	i := randomIter(*words)
	word = (*words)[i]
	u.SetIterWord(i)
	db.UpdateUser(u)
	return
}

func randomIter(a []string) int {
	size := len(a)
	r := rand.Intn(size)
	return r
}

func Word(chatID int64) (word string) {
	u := db.FindUser(chatID)
	words := storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	if u.GetIsToRu() {
		words = storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	}
	word = (*words)[u.GetIterWord()]
	return
}

func SetIsToRu(chatID int64, isToRu bool) {
	u := db.FindUser(chatID)
	u.SetIsToRu(isToRu)
	db.UpdateUser(u)
}

func SetLanguage(chatID int64, language string) {
	u := db.FindUser(chatID)
	u.SetLanguage(language)
	db.UpdateUser(u)
}

func Language(chaID int64) string {
	u := db.FindUser(chaID)
	return u.GetLanguage()
}

func SetTopic(chatID int64, topic string) {
	u := db.FindUser(chatID)
	u.SetTopic(topic)
	db.UpdateUser(u)
}

func ListWords(chatID int64) (list string) {
	u := db.FindUser(chatID)
	o := storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	t := storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	for i, w := range *o {
		if len(list) >= 350 {
			break
		}
		list += w + " - " + (*t)[i] + "\n"
	}
	return
}
