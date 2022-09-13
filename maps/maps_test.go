package main

import "testing"

func TestMap(t *testing.T) {
	dictionary := Dictionary{"test key": "test value"}

	got := dictionary.Search("test key")
	want := "test value"

	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test key")
	}
}
