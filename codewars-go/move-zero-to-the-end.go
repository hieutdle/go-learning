package kata

// MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})
// returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

func MoveZeros(arr []int) []int {
	res := make([]int, len(arr))
	ind := 0
	for i, v := range arr {
		if v != 0 {
			res[ind] = arr[i]
			ind++
		}
	}
	return res
}
