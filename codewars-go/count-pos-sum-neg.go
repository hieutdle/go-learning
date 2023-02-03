package kata

// First element is the count of positives numbers and the second element is sum of negative numbers.
// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].

func CountPositivesSumNegatives(numbers []int) []int {
	var res = []int{0, 0}

	for _, v := range numbers {
		switch {
		case v > 0:
			res[0]++
		case v < 0:
			res[1] += v
		}
	}
	return res
}
