package coin

import (
	"fmt"
	"sort"
)

func createCoinCache(amount int) []int {
	// create a slice from 0..amount + 1
	cache := make([]int, amount+1)

	// initialize the best answer (minimum number of coins)
	// to arbitrary value that is larger than the amount
	// in this case amount + 1 is fine
	for index, _ := range cache {
		cache[index] = amount + 1
	}

	return cache
}

// interative dynamic programming solution
func CoinChangeMin(coins []int, amount int) int {
	counter := 0
	ascendingCoins := sortAscendingImmutable(coins)
	// store the best answer (minimum number of coins) for each value
	// cache answer to each sub-problem in the cache

	// index in the cache corresponds to the amount sub-problem
	// the value in the cache respresents that sub-problem's best answer
	//   (minimum number of coins needed to make change)
	cache := createCoinCache(amount)

	// loop through the sub-problems
	// for each subproblem amount n,
	// find the minimum number of coins needed
	for amountSubProblem, _ := range cache {

		// set the answer of sub problem zero to 0
		// (the first value in the cache)
		if amountSubProblem == 0 {
			cache[amountSubProblem] = 0
			continue
		}

		// inside of each sub-problem, select a coin
		// and calculate difference from the sub-problem amount
		for _, coinSelected := range ascendingCoins {
			counter += 1
			currentBestAnswer := cache[amountSubProblem]
			remaining := amountSubProblem - coinSelected

			//fmt.Printf("currentMinimumCoins: %v, coinSelected: %v, remaining: %v\n", currentMinimumCoins, coinSelected, remaining)
			if remaining >= 0 {
				newMinimum := cache[remaining] + 1
				if newMinimum < currentBestAnswer {
					//fmt.Printf("updating cache at %v with new value: %v\n", amountSubProblem, newMinimum)
					cache[amountSubProblem] = newMinimum
				}
			} else {
				//fmt.Println("coin selected", coinSelected, "too large, skipping")
			}
		}
	}

	var finalAnswer = cache[len(cache)-1]

	if finalAnswer == amount+1 {
		finalAnswer = -1
	}

	fmt.Printf("final cache value: %v\n", finalAnswer)
	fmt.Printf("interated %v times.\n", counter)
	// the final value in the cache should be the optimal answer
	return finalAnswer
}

func sortAscendingImmutable(arr []int) []int {
	var copied = make([]int, len(arr))
	copy(copied, arr)
	sort.Sort(sort.IntSlice(copied))
	return copied
}
