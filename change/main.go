package main

import (
	"fmt"
	"saylerb/go-with-tests/change/coin"
)

func main() {
	coins := []coin.Coin{
		coin.Coin{Name: "penny", Value: 1},
		coin.Coin{Name: "2 cent", Value: 2},
		coin.Coin{Name: "nickel", Value: 5},
	}

	result := coin.CoinChange(coins, 11)

	fmt.Println(result)
}
