package hanoi

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("can create a stack with size 0", func(t *testing.T) {
		stack := Stack{}
		got := stack.Size()
		want := 0

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("can push a value onto the stack", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)
		want := Stack{1}

		if !reflect.DeepEqual(stack, want) {
			t.Errorf("got %v, wanted %v", stack, want)
		}
	})

	t.Run("can pop a value off of the stack", func(t *testing.T) {
		stack := Stack{1, 2}
		popped := stack.Pop()
		wantedPopped := 2
		want := Stack{1}

		if popped != 2 {
			t.Errorf("got %v, wanted %v", popped, wantedPopped)
		}
		if !reflect.DeepEqual(stack, want) {
			t.Errorf("got %v, wanted %v", stack, want)
		}
	})
}

func TestHanoi(t *testing.T) {

	t.Run("handle the base case of no disks to move", func(t *testing.T) {
		board := Board{A: &Stack{}, B: &Stack{}, C: &Stack{}}
		expected := Board{A: &Stack{}, B: &Stack{}, C: &Stack{}}

		board.Solve()

		if !reflect.DeepEqual(board, expected) {
			t.Errorf("got %v, wanted %v", board, expected)
		}
	})
	t.Run("single disk is moved from stack A to stack C", func(t *testing.T) {
		board := Board{A: &Stack{1}, B: &Stack{}, C: &Stack{}}

		board.Solve()
		got := board.C
		want := &Stack{1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("move two disks from stack A to stack C", func(t *testing.T) {
		board := Board{A: &Stack{2, 1}, B: &Stack{}, C: &Stack{}}

		board.Solve()
		got := board.C
		want := &Stack{2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("move three disks from stack A to stack C", func(t *testing.T) {
		board := Board{A: &Stack{3, 2, 1}, B: &Stack{}, C: &Stack{}}

		board.Solve()
		got := board.C
		want := &Stack{3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("3 pegs is 7 total moves", func(t *testing.T) {
		board := Board{A: &Stack{3, 2, 1}, B: &Stack{}, C: &Stack{}}

		board.Solve()

		got := board.TotalMoves()
		want := 7

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
