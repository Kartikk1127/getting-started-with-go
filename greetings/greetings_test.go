package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Kartikey"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) {
		t.Errorf(`Hello("Kartikey" = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
