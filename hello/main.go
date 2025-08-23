package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kartikey")
	fmt.Println(message)
}
