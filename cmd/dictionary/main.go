package main

import (
	"fmt"
	"time"

	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/telegram"
)

func main() {
	dict := storage.ReadDictionary(`dictionaries/english`)
	start := time.Now()
	s := telegram.GetRandomKey(dict)
	end := time.Now()
	fmt.Println(s)
	fmt.Println(end.Sub(start).Nanoseconds())
}
