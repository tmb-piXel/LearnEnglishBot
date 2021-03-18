package telegram

import "strings"

//Check if a element in slice
func contains(m map[int64]bool, e int64) bool {
	for a := range m {
		if a == e {
			return true
		}
	}
	return false
}

//Get random key from map
func getRandomKey(m map[string]string) string {
	for enW := range m {
		return enW
	}
	return ""
}

//Compaires the user's answer to the correct answer
func compaire(correct string, answer string) bool {
	flag := false
	answer = strings.ToLower(answer)
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
