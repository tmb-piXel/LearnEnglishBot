package main

import (
	"fmt"
	"time"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/telegram"
)

func main() {
	dict := storage.ReadDictionaries(`dictionaries`)
	start := time.Now()
	s := telegram.GetRandomWord(dict["english"])
	end := time.Now()
	fmt.Println(s)
	fmt.Println(end.Sub(start).Nanoseconds())
}
