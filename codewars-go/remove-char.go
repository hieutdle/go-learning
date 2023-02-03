package kata

// removes the first and last characters of a string

func RemoveChar(word string) string {
	return word[1:(len(word) - 1)]
}
