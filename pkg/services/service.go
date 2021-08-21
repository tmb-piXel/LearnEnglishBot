package services

import (
	"math/rand"

	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
	r "github.com/tmb-piXel/LearnEnglishBot/pkg/repositories"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
)

func NewUser(chatID int64) {
	u := m.NewUser(chatID)
	r.SaveUser(u)
}

func IsUserExist(chatID int64) bool {
	return r.IsUserExist(chatID)
}

func NewWord(chatID int64) (word string) {
	u := r.GetUser(chatID)
	words := u.GetTransleted()
	if u.GetIsToRu() {
		words = u.GetOriginal()
	}
	i := getR(*words)
	word = (*words)[i]
	u.SetIterWord(i)
	return
}

func Word(chatID int64) (word string) {
	u := r.GetUser(chatID)
	words := u.GetOriginal()
	if u.GetIsToRu() {
		words = u.GetTransleted()
	}
	word = (*words)[u.GetIterWord()]
	return
}

func getR(a []string) int {
	size := len(a)
	r := rand.Intn(size)
	return r
}

func SetIsToRu(chatID int64, isToRu bool) {
	u := r.GetUser(chatID)
	u.SetIsToRu(isToRu)
}

func SetLanguage(chatID int64, language string) {
	u := r.GetUser(chatID)
	u.SetLanguage(language)
}

func Language(chaID int64) string {
	u := r.GetUser(chaID)
	return u.GetLanguage()
}

func SetTopic(chatID int64, topic string) {
	u := r.GetUser(chatID)
	u.SetTopic(topic)
	u.SetOriginal(storage.GetOriginalWords(u.GetLanguage(), u.GetTopic()))
	u.SetTransleted(storage.GetTransletedWords(u.GetLanguage(), u.GetTopic()))
}

func ListWords(chatID int64) (list string) {
	u := r.GetUser(chatID)
	for i, w := range *u.GetOriginal() {
		list += w + " - " + (*u.GetTransleted())[i] + "\n"
	}
	return
}
