package kata

// returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
// n = 0  ==> [1]        # [2^0]
// n = 1  ==> [1, 2]     # [2^0, 2^1]
// n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]

func PowersOfTwo(n int) []uint64 {
	result := []uint64{1}
	for i := 0; i < n; i++ {
		result = append(result, result[i]*2)
	}
	return result
}
