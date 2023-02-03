package kata

//  Given n, populate an array with all numbers up to and including that number
// 10 --> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// 1 --> [1]

func monkeyCount(n int) []int {
	result := make([]int, n)

	for i := range result {
		result[i] = i + 1
	}

	return result
}
