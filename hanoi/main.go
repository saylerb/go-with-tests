package main

import (
	"saylerb/go-with-tests/hanoi/hanoi"
)

func main() {
	disks := 3

	init := hanoi.Stack{}
	for i := disks; i > 0; i-- {
		init.Push(i)
	}

	//fmt.Println("intialStack", init)
	board := hanoi.Board{A: &init, B: &hanoi.Stack{}, C: &hanoi.Stack{}}

	board.Solve()
}
