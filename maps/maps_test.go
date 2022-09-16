package main

import "testing"

func TestSearchDictionary(t *testing.T) {
	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test key": "test value"}
		got, _ := dictionary.Search("test key")
		want := "test value"

		assertString(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test key": "test value"}
		_, err := dictionary.Search("unknown key")

		assertError(t, err, ErrNotFound)
	})
}

func TestAddToDictionary(t *testing.T) {
	t.Run("add a new word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test key", "test value")

		got, _ := dictionary.Search("test key")
		want := "test value"

		assertString(t, got, want)
	})

	t.Run("add a word that already exists in dictionary", func(t *testing.T) {
		word := "test key"
		definition := "test value"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "another definition")

		found, _ := dictionary.Search("test key")

		assertError(t, err, ErrWordExists)
		assertString(t, found, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing word", func(t *testing.T) {
		dictionary := Dictionary{"test key": "test value"}

		dictionary.Update("test key", "updated value")

		found, _ := dictionary.Search("test key")

		assertString(t, found, "updated value")
	})

	t.Run("trying to update a word that does not exist results in an error", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("key that does not exist", "updated value")

		assertError(t, err, ErrNotFound)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test key")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("we expected an error to be returned")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
