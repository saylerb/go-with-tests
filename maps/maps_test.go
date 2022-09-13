package main

import "testing"

func TestMap(t *testing.T) {
	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test key": "test value"}
		got, _ := dictionary.Search("test key")
		want := "test value"

		if got != want {
			t.Errorf("got %q, want %q, given %q", got, want, "test key")
		}
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test key": "test value"}
		_, err := dictionary.Search("unknown key")
		want := "could not find the word you're looking for"

		if err == nil {
			t.Fatal("we expected an error to be returned")
		}

		if err.Error() != want {
			t.Errorf("got %q, want %q, given %q", err, want, "unknown key")
		}
	})
}
