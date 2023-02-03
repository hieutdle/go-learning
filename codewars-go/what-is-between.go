package kata

// return an array of all integers between the input parameters, including them.
// a = 1
// b = 4
// --> [1, 2, 3, 4]

func Between(a, b int) []int {
	var result []int
	for a <= b {
		result = append(result, a)
		a++
	}
	return result
}
