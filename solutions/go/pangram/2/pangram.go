package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for r := 'a'; r <= 'z'; r++ {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}
	return true
}
