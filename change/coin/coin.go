package coin

import (
	"fmt"
	"sort"
	"strings"
)

type Coin struct {
	Name  string
	Value int
}

type Change map[Coin]int

func (c Change) String() string {
	summary := []string{}

	// keys in map are not ordered, need to sort em
	var coins []Coin = []Coin{}
	for key, _ := range c {
		coins = append(coins, key)
	}
	decending := sortDecendingLambda(coins)

	for _, coin := range decending {
		amount := c[coin]
		printed := fmt.Sprintf("%v %v", amount, coin.Name)
		summary = append(summary, printed)
	}

	return strings.Join(summary, ",\n")
}

func CoinChange(coins []Coin, amount int) string {
	if len(coins) == 0 {
		return "no change can be made, no coins available :("
	}
	if amount == 0 {
		return "no change can be made for amount 0"
	}
	decending := sortDecendingLambda(coins)

	// create map here for the change object based on the
	// available coins
	var result Change = make(Change)

	for _, coin := range decending {
		for amount >= coin.Value {
			result[coin] += 1
			amount -= coin.Value
		}
	}
	return result.String()
}

func sortDecending(arr []int) []int {
	var copied = make([]int, len(arr))
	copy(copied, arr)                                 // copy the contents to a new array
	var slice sort.IntSlice = sort.IntSlice(copied)   // wrap in IntSlice type
	var reversed sort.Interface = sort.Reverse(slice) // return new sort Interface with reversed Less method
	sort.Sort(reversed)
	return copied
}

func sortDecendingMut(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

func sortDecendingLambda(arr []Coin) []Coin {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value > arr[j].Value // reverse sort
	})
	return arr
}
