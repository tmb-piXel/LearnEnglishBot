package storage

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Read the dictionary from the file
func ReadDictionaries(pathDictionary string) (dictionaries map[string]map[string]string) {
	files, err := ioutil.ReadDir(pathDictionary)
	if err != nil {
		log.Printf("Failed read path with dictionaries: %s", err)
	}

	dictionaries = make(map[string]map[string]string)

	for _, f := range files {
		file, err := os.Open(pathDictionary + "/" + f.Name())

		if err != nil {
			log.Printf("Failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		dictionary := make(map[string]string)

		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, "-")
			original := string(words[0])
			translated := string(words[1])
			dictionary[original] = translated
		}

		file.Close()
		dictionaries[f.Name()] = dictionary
	}

	return
}
