package kata

// reverses the string passed into it.
// 'world'  =>  'dlrow'
// 'word'   =>  'drow'

func Solution(word string) string {
	reversed := ""
	for _, v := range word {
		reversed = string(v) + reversed
	}
	return reversed
}
