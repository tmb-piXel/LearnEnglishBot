package main

import (
	"fmt"

	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
)

func main() {
	dict := storage.ReadDictionary(`dictionary`)
	fmt.Println(dict)
}
