package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a specific person when provided", func(t *testing.T) {
		got := Hello("Brian")
		want := "Hello, Brian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when no name is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tell go that this function is a helper so that line number of failure
	// comes from caller and not the helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
