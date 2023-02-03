package kata

import "fmt"

//  Jenny is in love with Johnny, and would like to greet him slightly different.

func Greet(name string) string {
	if name == "Johnny" {
		name = "my love"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
