package twelve

import (
	"fmt"
	"strings"
)

type verse struct {
	day   string
	words string
}

var verseBits = map[int]verse{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func composeLastPart(i int) string {
	if i == 1 {
		return verseBits[i].words
	}
	upToLast := ""
	for j := i; j > 1; j-- {
		upToLast += verseBits[j].words + ", "
	}
	return fmt.Sprintf("%sand %s", upToLast, verseBits[1].words)
}

func Verse(i int) string {
	base := "On the %s day of Christmas my true love gave to me: %s."
	return fmt.Sprintf(base, verseBits[i].day, composeLastPart(i))
}

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return strings.TrimSpace(song)
}
