package kata

import "strconv"

/* Given an array of integers as strings and numbers,
return the sum of the array values as if all were numbers.
Return your answer as a number.
{9, 3, "7", "3"} => 22)
*/

func SumMix(arr []any) int {

	sum := 0

	for _, v := range arr {
		switch v.(type) {
		case int:
			sum += v.(int)
		case string:
			s, _ := strconv.Atoi(v.(string))
			sum += s
		}
	}
	return sum
}
