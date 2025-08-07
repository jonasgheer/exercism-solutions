package robotname

import "math/rand"

type Robot struct {
	name string
}

var names []string

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = generateName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	for {
		name := string([]rune{randLetter(), randLetter(), randDigit(), randDigit(), randDigit()})
		if contains(names, name) {
			continue
		}
		names = append(names, name)
		return name
	}
}

func contains(names []string, name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}

func randLetter() rune {
	return letters[rand.Intn(len(letters))]
}

func randDigit() rune {
	return digits[rand.Intn(len(digits))]
}
