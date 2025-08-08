package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	result := ""
	count := 0
	for i := 0; i < len(input); i++ {
		count++
		if i == len(input)-1 || input[i] != input[i+1] {
			if count > 1 {
				result += strconv.Itoa(count)
			}
			result += string(input[i])
			count = 0
		}
	}
	return result

}
func RunLengthDecode(input string) string {
	result := ""
	currPos := 0
	for {
		if currPos >= len(input) {
			break
		}
		if isDigit(input[currPos]) {
			pos := currPos
			for isDigit(input[currPos]) {
				currPos++
			}
			n, err := strconv.Atoi(input[pos:currPos])
			if err != nil {
				panic("atoi failed?")
			}
			result += strings.Repeat(string(input[currPos]), n)
		} else {
			result += string(input[currPos])
		}

		currPos++
	}
	return result
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
