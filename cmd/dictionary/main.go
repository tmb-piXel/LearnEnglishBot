package main

import (
	"fmt"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
)

func main() {
	pathDictionary := "dictionaries"
	storage.InitDictionaries(pathDictionary)

	fmt.Println(storage.GetLanguages())
	fmt.Println(storage.GetCode("Spain"))
	fmt.Println(storage.GetTopicTitles("German"))
	fmt.Println(storage.GetOriginalWords("English", "allE"))
	fmt.Println(storage.GetTransletedWords("English", "allE"))
}
