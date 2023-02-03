package kata

import "strings"

// remove the spaces from the string, then return the resultant string.

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
