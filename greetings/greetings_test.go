package greetings

import "testing"

func TestHello(t *testing.T) {
	name := "Kartikey"
	message := "Hi, Kartikey. Welcome!"
	got := Hello(name)

	if got != message {
		t.Errorf("Hello(%q) = %q, want %q", name, message, got)
	}
	t.Log("Testing finished successfully")
}
