package kata

// Given an array of integers your solution should find the smallest integer.

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
