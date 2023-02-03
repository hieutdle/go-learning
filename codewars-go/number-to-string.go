package kata

import "strconv"

// "1234" --> 1234
// "605"  --> 605
// "1405" --> 1405
// "-7" --> -7

func NumberToString(n int) string {
	return strconv.Itoa(n)
}
