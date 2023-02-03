package kata

import "strings"

// split a string and convert it into an array of words.

// "Robin Singh" ==> ["Robin", "Singh"]
//"I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

func StringToArray(str string) []string {
	return strings.Fields(str)
}
