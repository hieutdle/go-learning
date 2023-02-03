package kata

//  return a new array with each value doubled.
// [1, 2, 3] --> [2, 4, 6]

func Maps(x []int) []int {
	doubled := make([]int, len(x))
	for i := range x {
		doubled[i] = x[i] * 2
	}
	return doubled
}
