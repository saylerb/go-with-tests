package coin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCoinResult(t *testing.T) {
	t.Run("making sure we can print out the Change object", func(t *testing.T) {
		change := Change{Coin{name: "penny", value: 1}: 2}
		got := fmt.Sprint(change)

		want := "2 penny"

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestCoin(t *testing.T) {
	t.Run("making change zero returns zero", func(t *testing.T) {
		penny := Coin{name: "penny", value: 1}
		got := CoinChange([]Coin{penny}, 0)
		want := "no change can be made for amount 0"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("making change for 10 cents with a single dime", func(t *testing.T) {
		dime := Coin{name: "dime", value: 10}

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
		oneCent := Coin{name: "one cent coin", value: 1}
		twoCent := Coin{name: "two cent coin", value: 2}
		fiveCent := Coin{name: "five cent coin", value: 5}
		got := CoinChange([]Coin{oneCent, twoCent, fiveCent}, 11)
		want := "2 five cent coin,\n1 one cent coin"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("making change for 177 cents with 1, 5, 10, 50, 100, 500 cent coins", func(t *testing.T) {
		availableCoins := []Coin{
			Coin{name: "one yen coin", value: 1},
			Coin{name: "five yen coin", value: 5},
			Coin{name: "ten yen coin", value: 10},
			Coin{name: "fifty yen coin", value: 50},
			Coin{name: "hundred yen coin", value: 100},
			Coin{name: "five hundred yen coin", value: 500},
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
		oneCent := Coin{name: "one cent coin", value: 1}
		twoCent := Coin{name: "two cent coin", value: 2}
		fiveCent := Coin{name: "five cent coin", value: 5}

		got := sortDecendingLambda([]Coin{oneCent, twoCent, fiveCent})
		want := []Coin{fiveCent, twoCent, oneCent}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
