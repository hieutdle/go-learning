package kata

// Returns "Even" for even numbers or "Odd" for odd numbers

func EvenOrOdd(number int) string {
	str := "Odd"
	if number%2 == 0 {
		str = "Even"
	}
	return str
}
