package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountdownOperations struct {
	Calls []string
}

func (s *CountdownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *CountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, string(p))
	return
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{sleep: time.Sleep, duration: 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
