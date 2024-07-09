package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("should change language depending on 2nd parameter", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("should return french version if French is the second parameter", func(t *testing.T) {
		got := Hello("Josh", "French")
		want := "Bonjour, Josh!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("should return italian version if Italian is the second parameter", func(t *testing.T) {
		got := Hello("Josh", "Italian")
		want := "Ciao, Josh!"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	got := Hello("Josh", "French")
	fmt.Println(got)
	// Output: Bonjour, Josh!
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
