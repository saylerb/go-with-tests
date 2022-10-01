package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprintln(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
