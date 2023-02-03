package kata

import "strconv"

// Given a non-negative integer, 3
// Return a string with a murmur: "1 sheep...2 sheep...3 sheep..."

func countSheep(num int) string {
	str := ""
	for i := 1; i <= num; i++ {
		str += strconv.Itoa(i) + " sheep..."
		// fmt.Sprintf("%d sheep...", i)
	}
	return str
}
