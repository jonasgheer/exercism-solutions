package raindrops

import (
	"fmt"
	"strings"
)

// Convert returns the corresponding sound of n
func Convert(n int) string {
	var sb strings.Builder
	if n%3 == 0 {
		sb.WriteString("Pling")
	}
	if n%5 == 0 {
		sb.WriteString("Plang")
	}
	if n%7 == 0 {
		sb.WriteString("Plong")
	}
	sound := sb.String()
	if sound == "" {
		return fmt.Sprint(n)
	}
	return sound
}
