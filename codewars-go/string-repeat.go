package kata

import "strings"

// accepts an integer n and a string s as parameters
// returns a string of s repeated exactly n times.
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}
