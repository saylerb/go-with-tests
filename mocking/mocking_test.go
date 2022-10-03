package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("test that the output is correct", func(t *testing.T) {
		var buffer *bytes.Buffer = &bytes.Buffer{}
		var spy Sleeper = &SpySleeper{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("test the calls to sleep and write are in the correct order", func(t *testing.T) {
		var spy *CountdownOperations = &CountdownOperations{}

		Countdown(spy, spy)

		wantedCalls := []string{
			"3\n",
			"sleep",
			"2\n",
			"sleep",
			"1\n",
			"sleep",
			"Go!\n",
		}

		if !reflect.DeepEqual(wantedCalls, spy.Calls) {
			t.Errorf("incorrect ordering of calls, wanted %q, got %q", wantedCalls, spy.Calls)
		}
	})

	t.Run("test the configuration of the sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}

	})
}
