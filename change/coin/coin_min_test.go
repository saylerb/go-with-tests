package coin

import (
	"reflect"
	"testing"
)

func TestCoinMin(t *testing.T) {
	t.Run("making change zero returns zero", func(t *testing.T) {
		got := CoinChangeMin([]int{1}, 0)
		want := 0

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 10 cents with a single dime", func(t *testing.T) {
		got := CoinChangeMin([]int{10}, 10)
		want := 1

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("when no available coins to make change with, return -1", func(t *testing.T) {
		got := CoinChangeMin([]int{}, 5)
		want := -1

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 11 cents with 1, 2 and 5 cent coins", func(t *testing.T) {
		got := CoinChangeMin([]int{1, 2, 5}, 11)
		want := 3

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 177 cents with 1, 5, 10, 50, 100, 500 cent coins", func(t *testing.T) {
		got := CoinChangeMin([]int{1, 5, 10, 50, 100, 500}, 177)
		// 100, 50, 10, 10, 5,1,1
		want := 7

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("if you can make change with fewer numbers of smaller denomenations", func(t *testing.T) {
		got := CoinChangeMin([]int{1, 4, 5}, 8)
		// want 4,4
		// instead of 5, 1, 1, 1

		want := 2
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("min coins to make 55 cents with 1,20,25,and 50 cent is 4", func(t *testing.T) {
		// coins: 1, 10, 25, 50
		// amount: 55 cents

		// Dynamic: 25 + 10 + 10 + 10 (4 total)
		// Greedy: 50 + 1 + 1 + 1 + 1 + 1 (6 total)

		got := CoinChangeMin([]int{1, 10, 25, 50}, 55)

		want := 4
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("[186,419,83,408]", func(t *testing.T) {
		got := CoinChangeMin([]int{186, 419, 83, 408}, 6249)

		want := 20
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestCoinCache(t *testing.T) {
	got := createCoinCache(5)
	want := []int{6, 6, 6, 6, 6, 6}
	//            0  1  2  3  4  5
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
