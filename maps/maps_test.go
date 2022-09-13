package main

import "testing"

func TestMap(t *testing.T) {
	dictionary := map[string]string{"test key": "test value"}

	got := Search(dictionary, "test key")
	want := "test value"

	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test key")
	}
}
