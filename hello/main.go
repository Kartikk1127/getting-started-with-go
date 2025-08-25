package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// set the properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
