package kata

// you are given a number and have to make it negative.
// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0

func MakeNegative(x int) int {
	if x > 0 {
		x = x * -1
	}
	return x
}
