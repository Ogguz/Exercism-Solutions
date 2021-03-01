package twofer

import "fmt"

// ShareWith gets the name and returns a string
func ShareWith(name string) string {

	if name != "" {
		return fmt.Sprintf("One for %s, one for me.",name)
	}
	return "One for you, one for me."
}
