package kata

// Build a function that returns an array of integers from n to 1 where n>0.
// n=5 --> [5,4,3,2,1]

func ReverseSeq(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = n - i
	}
	return seq
}
