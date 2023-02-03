package kata

/* Deadfish has 4 commands, each 1 character long:
i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Parse("iiisdoso") == []int{8, 64} */

func Parse(data string) []int {
	df := []int{}
	value := 0
	for _, v := range data {
		switch v {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value = value * value
		case 'o':
			df = append(df, value)
		}
	}
	return df
}
