package main

import "testing"

func TestSearchDictionary(t *testing.T) {

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

		if err == nil {
			t.Fatal("we expected an error to be returned")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAddToDictionary(t *testing.T) {
	t.Run("add a key to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test key", "test value")

		got, _ := dictionary.Search("test key")
		want := "test value"

		if got != want {
			t.Errorf("got %q, want %q, given %q", got, want, "test key")
		}
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
