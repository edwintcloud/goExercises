package isogram

import "strings"

// IsIsogram checks if a word is an isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters := make(map[string]int)

	// iterate through letters in word
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		letters[c]++
		if letters[c] > 1 && c != "-" && c != " " {
			return false
		}
	}

	// if we made it here, return true
	return true
}
