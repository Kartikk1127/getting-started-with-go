package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello the function name is in capital letter so it could be called from outside this package
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	//message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	//return a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
