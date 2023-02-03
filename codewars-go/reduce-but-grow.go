package kata

//  return the result of multiplying the values together in order.
// [1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

func Grow(arr []int) int {
	res := 1
	for _, v := range arr {
		res *= v
	}
	return res
}
