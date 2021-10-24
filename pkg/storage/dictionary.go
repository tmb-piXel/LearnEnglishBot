package storage

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type dictionary struct {
	language string
	topics   map[string]topic
}

type topic struct {
	title           string
	originalwords   []string
	translatedWords []string
}

var dictionaries = make(map[string]dictionary)

func InitDictionaries(path string) {
	dictionaries = readDictionaries(path)
}

func readDictionaries(path string) (dictionaries map[string]dictionary) {
	dictionaries = make(map[string]dictionary)
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Failed read path with dictionaries: %s", err)
	}
	for _, d := range dirs {
		files, err := ioutil.ReadDir(path + "/" + d.Name())
		if err != nil {
			log.Printf("Failed read dir %s: %s", d.Name(), err)
		}
		topics := make(map[string]topic)
		for _, f := range files {
			titleTopic := f.Name()
			topic := topic{title: titleTopic}
			f, err := os.Open((path + "/" + d.Name() + "/" + f.Name()))
			if err != nil {
				log.Printf("Failed read topic %s: %s", titleTopic, err)
			}
			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				words := strings.Split(line, "----")
				original := string(words[0])
				translated := string(words[1])
				topic.originalwords = append(topic.originalwords, original)
				topic.translatedWords = append(topic.translatedWords, translated)
			}
			topics[titleTopic] = topic
		}
		lang := d.Name()
		dictionaries[lang] = dictionary{language: lang, topics: topics}
	}
	return
}

func GetLanguages() (languages []string) {
	for _, d := range dictionaries {
		languages = append(languages, d.language)
	}
	return
}

func GetTopicTitles(language string) (titles []string) {
	for _, t := range dictionaries[language].topics {
		titles = append(titles, t.title)
	}
	return
}

func GetOriginalWords(language, topic string) (originalwords *[]string) {
	for _, t := range dictionaries[language].topics {
		if t.title == topic {
			originalwords = &t.originalwords
			break
		}
	}
	return
}

func GetTransletedWords(language, topic string) (translatedWords *[]string) {
	for _, t := range dictionaries[language].topics {
		if t.title == topic {
			translatedWords = &t.translatedWords
			break
		}
	}
	return
}
