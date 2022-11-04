package coin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCoinResult(t *testing.T) {
	t.Run("making sure we can print out the Change object", func(t *testing.T) {
		change := Change{Coin{Name: "penny", Value: 1}: 2}
		got := fmt.Sprint(change)

		want := "2 penny"

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestCoin(t *testing.T) {
	t.Run("making change zero returns zero", func(t *testing.T) {
		penny := Coin{Name: "penny", Value: 1}
		got := CoinChange([]Coin{penny}, 0)
		want := "no change can be made for amount 0"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("making change for 10 cents with a single dime", func(t *testing.T) {
		dime := Coin{Name: "dime", Value: 10}

		got := CoinChange([]Coin{dime}, 10)
		want := "1 dime"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("when no available coins to make change with, return -1", func(t *testing.T) {
		got := CoinChange([]Coin{}, 5)
		want :=
			"no change can be made, no coins available :("

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("making change for 11 cents with 1, 2 and 5 cent coins", func(t *testing.T) {
		oneCent := Coin{Name: "one cent coin", Value: 1}
		twoCent := Coin{Name: "two cent coin", Value: 2}
		fiveCent := Coin{Name: "five cent coin", Value: 5}
		got := CoinChange([]Coin{oneCent, twoCent, fiveCent}, 11)
		want := "2 five cent coin,\n1 one cent coin"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("making change for 177 cents with 1, 5, 10, 50, 100, 500 cent coins", func(t *testing.T) {
		availableCoins := []Coin{
			Coin{Name: "one yen coin", Value: 1},
			Coin{Name: "five yen coin", Value: 5},
			Coin{Name: "ten yen coin", Value: 10},
			Coin{Name: "fifty yen coin", Value: 50},
			Coin{Name: "hundred yen coin", Value: 100},
			Coin{Name: "five hundred yen coin", Value: 500},
		}

		got := CoinChange(availableCoins, 177)

		want := `1 hundred yen coin,
1 fifty yen coin,
2 ten yen coin,
1 five yen coin,
2 one yen coin`

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
func TestDecending(t *testing.T) {
	t.Run("test can sort ints decending", func(t *testing.T) {
		got := sortDecending([]int{1, 2, 5})
		want := []int{5, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("test can sort ints decending mutating the original array", func(t *testing.T) {
		got := sortDecendingMut([]int{1, 2, 5})
		want := []int{5, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("test can sort ints with a lambda", func(t *testing.T) {
		oneCent := Coin{Name: "one cent coin", Value: 1}
		twoCent := Coin{Name: "two cent coin", Value: 2}
		fiveCent := Coin{Name: "five cent coin", Value: 5}

		got := sortDecendingLambda([]Coin{oneCent, twoCent, fiveCent})
		want := []Coin{fiveCent, twoCent, oneCent}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
