package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet(&buffer, "Brian")

	got := buffer.String()
	want := "Hello, Brian"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

