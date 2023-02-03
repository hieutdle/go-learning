package kata

// the summation of every number from 1 to num.
// 2 -> 3 (1 + 2)
// 8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)

func Summation(n int) int {
	return n * (n + 1) / 2
}
