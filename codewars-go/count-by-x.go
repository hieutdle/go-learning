package kata

// return an array of the first n multiples of x.

// countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
// countBy(2,5)  should return {2,4,6,8,10}

func CountBy(x, n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, x*i)
	}
	return result
}
