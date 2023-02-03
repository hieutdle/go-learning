package kata

// Takes a boolean value
// Return a "Yes" string for true, or a "No" string for false.

func BoolToWord(word bool) string {
	str := ""
	switch word {
	case true:
		str = "Yes"
	case false:
		str = "No"
	}
	return str
}
