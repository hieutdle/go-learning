package kata

// Replace any digit below 5 with '0' and any digit 5 and above with '1'.

func FakeBin(x string) string {
	result := ""
	for _, char := range x {
		if char > 52 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}
