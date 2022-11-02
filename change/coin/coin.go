package coin

import "sort"

func CoinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	decending := sortDecending(coins)

	var result int

	for _, coin := range decending {
		for amount >= coin {
			result += 1
			amount -= coin
		}
	}
	return result
}

func sortDecending(arr []int) []int {
	var copied = make([]int, len(arr))
	copy(copied, arr)                                 // copy the contents to a new array
	var slice sort.IntSlice = sort.IntSlice(copied)   // wrap in IntSlice type
	var reversed sort.Interface = sort.Reverse(slice) // return new sort Interface with reversed Less method
	sort.Sort(reversed)
	return copied
}
