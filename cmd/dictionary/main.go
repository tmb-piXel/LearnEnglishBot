package main

import (
	"fmt"

	"github.com/tmb-piXel/telegramBotForLearningEnglish/pkg/storage"
)

func main() {
	fmt.Println(storage.ReadDictionary())
}
