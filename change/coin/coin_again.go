package coin

import (
	"fmt"
	"sort"
)

func CoinChangeMinXXX(coins []int, amount int) int {
	totalCoins := len(coins)
	var changePossibilities []int

	ascendingCoins := sortAscendingImmutable(coins)

	for index, _ := range ascendingCoins {
		availableCoins := ascendingCoins[:totalCoins-index]

		fmt.Println("availableCoins", availableCoins)
		result := changeHelper(availableCoins, amount)
		changePossibilities = append(changePossibilities, result)
	}

	fmt.Println("possibilities:", changePossibilities)

	sorted := sortDecendingImmutable(changePossibilities)
	return sorted[len(sorted)-1]
}

func changeHelper(coins []int, amount int) int {
	var result []int
	if len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	decending := sortDecendingImmutable(coins)

	for _, coin := range decending {
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}
	totalCoins := len(result)
	fmt.Printf("total: %v, coins: %v, remaining: %v\n", totalCoins, result, amount)
	return len(result)

}

func sortDecendingImmutable(arr []int) []int {
	var copied = make([]int, len(arr))
	copy(copied, arr)
	var slice sort.IntSlice = sort.IntSlice(copied)
	var reversed sort.Interface = sort.Reverse(slice)
	sort.Sort(reversed)
	return copied
}
