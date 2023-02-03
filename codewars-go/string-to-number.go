package kata

import "strconv"

// transform a string into a number
// "1234" --> 1234
// "605"  --> 605

func StringToNumber(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
