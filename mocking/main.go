package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(writer, fmt.Sprintf("%d\n", i))
	}
	fmt.Fprint(writer, "Go!\n")
}

func main() {
	Countdown(os.Stdout)
}
