package kata

//  a function that counts the number of sheep present in the array (true means present).
/*
[true,  true,  true,  false,
 true,  true,  true,  true ]
*/

func CountSheeps(numbers []bool) int {
	count := 0
	for i := range numbers {
		if numbers[i] {
			count++
		}
	}
	return count
}
