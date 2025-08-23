package greetings

import "fmt"

// Hello the function name is in capital letter so it could be called from outside this package
func Hello(name string) string {
	// return a greeting that embeds name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
