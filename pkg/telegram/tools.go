package telegram

import (
	"strings"
)

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
