package coin

import (
	"reflect"
	"testing"
)

func TestCoin(t *testing.T) {
	t.Run("making change zero returns zero", func(t *testing.T) {
		got := CoinChange([]int{1}, 0)
		want := 0

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 10 cents with a single dime", func(t *testing.T) {
		got := CoinChange([]int{10}, 10)
		want := 1

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("when no available coins to make change with, return -1", func(t *testing.T) {
		got := CoinChange([]int{}, 5)
		want := -1

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 11 cents with 1, 2 and 5 cent coins", func(t *testing.T) {
		got := CoinChange([]int{1, 2, 5}, 11)
		want := 3

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("making change for 177 cents with 1, 5, 10, 50, 100, 500 cent coins", func(t *testing.T) {
		got := CoinChange([]int{1, 5, 10, 50, 100, 500}, 177)
		// 100, 50, 10, 10, 5,1,1
		want := 7

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
func TestDecending(t *testing.T) {
	t.Run("test can sort the coins decending", func(t *testing.T) {
		got := sortDecending([]int{1, 2, 5})
		want := []int{5, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
