package scrabble

import "strings"

var letterScores = map[string]int{
	"aeioulnrst": 1,
	"dg":         2,
	"bcmp":       3,
	"fhvwy":      4,
	"k":          5,
	"jx":         8,
	"qz":         10,
}

func Score(word string) int {
	chars := strings.Split(strings.ToLower(word), "")
	totalScore := 0
	for _, c := range chars {
		for letters, score := range letterScores {
			if strings.Contains(letters, c) {
				totalScore += score
			}
		}
	}
	return totalScore
}
