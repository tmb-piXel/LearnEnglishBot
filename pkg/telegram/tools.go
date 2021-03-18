package telegram

//Check if a element in slice
func contains(m map[int64]bool, e int64) bool {
	for a := range m {
		if a == e {
			return true
		}
	}
	return false
}

func getRandomKey(m map[string]string) string {
	for enW := range m {
		return enW
	}
	return ""
}
