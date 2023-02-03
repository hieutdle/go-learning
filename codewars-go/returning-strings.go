package kata

import "fmt"

// return a greeting statement that uses an input.
// your program should return, "Hello, <name> how are you doing today?".

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
