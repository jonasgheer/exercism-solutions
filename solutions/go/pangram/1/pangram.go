package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	var r rune
	for r = 97; r < 123; r++ {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}
	return true
}
