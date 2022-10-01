package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

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
	Countdown(os.Stdout, &DefaultSleeper{})
}
