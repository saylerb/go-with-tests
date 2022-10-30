package hanoi

import "fmt"

type Stack []int

func (s Stack) Size() int {
	return len(s)
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	slice := []int(*s) // convert to slice type

	last := slice[len(slice)-1]
	slice = slice[:len(slice)-1]

	*s = Stack(slice)
	return last
}

type Board struct {
	A     *Stack
	B     *Stack
	C     *Stack
	moves int
}

func (b Board) String() string {
	return fmt.Sprintf("Board{A: %+v, B: %+v,C: %+v, moves: %d}", b.A, b.B, b.C, b.moves)
}

func (b *Board) Solve() {
	solveHanoi(b, len(*b.A), b.A, b.C, b.B)
}

func (b Board) TotalMoves() int {
	return b.moves
}

func solveHanoi(b *Board, totalDisks int, from *Stack, to *Stack, other *Stack) {
	if totalDisks == 0 {
		return
	}

	// solve subproblem of moving n-1 disks to the _other_ stack
	solveHanoi(b, totalDisks-1, from, other, to)

	// move exposed bottom disk to the desired destination
	bottomDisk := from.Pop()
	to.Push(bottomDisk)
	b.moves += 1
	fmt.Println(b)

	// solve subproblem of moving n-1 disks from _other_ stack to dest
	solveHanoi(b, totalDisks-1, other, to, from)
}
