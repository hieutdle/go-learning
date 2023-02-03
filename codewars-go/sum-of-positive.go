package kata

// return the sum of all of the positives ones.
// [1,-4,7,12] => 1 + 7 + 12 = 20

func PositiveSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return sum
}
