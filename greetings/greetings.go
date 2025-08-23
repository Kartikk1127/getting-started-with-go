package greetings

import (
	"errors"
	"fmt"
)

// Hello the function name is in capital letter so it could be called from outside this package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
