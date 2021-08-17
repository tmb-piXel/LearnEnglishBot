package telegram

import (
	"math/rand"
	"strings"
)

//Contains - check if a element in slice
func Contains(m map[int64]bool, e int64) bool {
	for a := range m {
		if a == e {
			return true
		}
	}
	return false
}

//GetRandomKey - get random key from map
func GetRandomWord(m map[string]string) string {
	for word := range m {
		return word
	}
	return ""
}

func GetRand(m map[string]string) string {
	for _, word := range m {
		return word
	}
	return ""
}

func GetR(a []string) int {
	size := len(a)
	r := rand.Intn(size)
	return r
}

//Compaire - compaires the user's answer to the correct answer
func CheckAnswer(correct string, answer string) bool {
	flag := false
	answer = strings.ToLower(answer)
	answer = strings.Trim(answer, " ")
	correct = strings.ToLower(correct)
	words := strings.Split(correct, "/")
	for _, word := range words {
		word = strings.Trim(word, " ")
		if answer == word {
			flag = true
		} else {
			continue
		}
	}
	return flag
}
