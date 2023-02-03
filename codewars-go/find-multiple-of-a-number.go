package kata

// a program that takes a value, integer
// returns a list of its multiples up to another value, limit
// if the parameters passed are(2, 6), the function should return [2, 4, 6]

func FindMultiples(integer, limit int) []int {
	nums := limit / integer
	var res []int
	count := 1
	for count <= nums {
		res = append(res, integer*count)
		count++
	}
	return res
}
