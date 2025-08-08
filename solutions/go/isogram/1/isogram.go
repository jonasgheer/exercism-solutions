package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters := make(map[string]int)
	for _, letter := range strings.Split(word, "") {
		if letter == "-" || letter == " " {
			continue
		}
		if _, exists := letters[letter]; exists {
			letters[letter] += 1
		} else {
			letters[letter] = 1
		}
		for _, count := range letters {
			if count > 1 {
				return false
			}
		}
	}
	return true
}
