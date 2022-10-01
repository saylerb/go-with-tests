package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
