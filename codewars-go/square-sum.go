package kata

// squares each number passed into it and then sums the results together.
// [1, 2, 2] should return 9 because 1^2 + 2^2 + 2^2 = 9.

func SquareSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v * v
	}
	return sum
}
