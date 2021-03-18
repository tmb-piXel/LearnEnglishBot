package storage

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//Read the dictionary from the file
func ReadDictionary(dictionaryFile string) (dictionary map[string]string) {
	file, err := os.Open(dictionaryFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	dictionary = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "-")
		englishWord := string(words[0])
		russianWord := string(words[1])
		dictionary[englishWord] = russianWord
	}

	file.Close()

	return dictionary
}
