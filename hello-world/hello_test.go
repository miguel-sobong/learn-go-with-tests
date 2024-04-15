package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hello to inputted name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should default to 'Hello, World' when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should say hello to inputted name in Bisaya", func(t *testing.T) {
		got := Hello("Chris", "bisaya")
		want := "Hoy, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should default to english when language is empty or invalid", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

		got = Hello("Chris", "INVALID")
		want = "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
