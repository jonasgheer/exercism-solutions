package atbash

import (
	"slices"
	"strings"
)

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var alphabetRev = []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}

func Atbash(s string) string {
	result := ""
	s = strings.ToLower(s)
	for _, c := range s {
		index := slices.Index(alphabet, c)
		if index == -1 {
			if isNumber(c) {
				result += string(c)
			}
		} else {
			result += string(alphabetRev[index])
		}
	}
	return chunk(result)
}

func chunk(s string) string {
	result := ""
	count := 0
	for _, c := range s {
		if count == 5 {
			result += " "
			count = 0
		}
		result += string(c)
		count++
	}
	return result
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
