package kata

// return the digits of this number within an array in reverse order.
// 35231 => [1,3,2,5,3]
// 0 => [0]

func Digitize(n int) []int {
	digits := []int{}
	for {
		digits = append(digits, n%10)
		n /= 10
		if n == 0 {
			return digits
		}
	}
}
