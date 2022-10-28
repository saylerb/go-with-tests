package main

import (
	"fmt"
	"sort"
)

type RomansMap map[int]string

var romans = RomansMap{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
	0:    "",
}

func main() {
	fmt.Println("romans")
}

func IntegerToRoman(number int) string {
	var current int

	numeral, present := romans[number]

	if present {
		return numeral
	}

	keys := make([]int, 0)

	for k, _ := range romans {
		keys = append(keys, k)
	}

	// sort the keys in map decending
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, key := range keys {
		if number >= key {
			current = key
			break
		}
	}

	return romans[current] + IntegerToRoman(number-current)
}
