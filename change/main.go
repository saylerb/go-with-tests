package main

import (
	"fmt"
	"saylerb/go-with-tests/change/coin"
)

func main() {
	result := coin.CoinChange([]int{1, 2, 5}, 11)

	fmt.Println(result)
}
