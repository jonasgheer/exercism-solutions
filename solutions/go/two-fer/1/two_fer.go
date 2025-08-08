package twofer

import "fmt"

// ShareWith returns message with name. Uses "you" if name is empty string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
