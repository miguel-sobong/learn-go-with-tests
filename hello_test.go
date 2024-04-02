package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hello to inputted name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should default to 'Hello, World' when name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Worsld"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
