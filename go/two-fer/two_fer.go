package twofer

import "fmt"

// ShareWith should have a comment documenting it. No it shouldn't
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
