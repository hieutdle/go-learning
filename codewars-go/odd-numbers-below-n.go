package kata

// return the number of positive odd numbers below n
// 7  -> 3 (because odd numbers below 7 are [1, 3, 5])
// 15 -> 7 (because odd numbers below 15 are [1, 3, 5, 7, 9, 11, 13])

func OddCount(n int) int {
	return n / 2
}
