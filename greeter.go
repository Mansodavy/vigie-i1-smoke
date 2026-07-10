package greeter

import "fmt"

// Greet returns a greeting for name.
func Greet(name string) string {
	return "Hello, " + name
}

func PrintGreeting(name string) {
	fmt.Println(Greet(name))
}
