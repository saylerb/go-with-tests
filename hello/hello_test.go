package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a specific person when provided", func(t *testing.T) {
		got := Hello("Brian")
		want := "Hello, Brian"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when no name is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
