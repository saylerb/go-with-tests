package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	var buffer *bytes.Buffer = &bytes.Buffer{}
	var spySleeper *SpySleeper = &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!
`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to the sleeper, want 3 go %d", spySleeper.Calls)
	}
}
