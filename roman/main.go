package main

import (
	"fmt"
	"sort"
	"strings"
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

func IntegerToRomanBrute(number int) string {
	if number >= 1000 {
		return "M" + IntegerToRomanBrute(number-1000)
	}
	if number >= 900 {
		return "CM" + IntegerToRomanBrute(number-900)
	}
	if number >= 500 {
		return "D" + IntegerToRomanBrute(number-500)
	}
	if number >= 400 {
		return "CD" + IntegerToRomanBrute(number-400)
	}
	if number >= 100 {
		return "C" + IntegerToRomanBrute(number-100)
	}
	if number >= 90 {
		return "XC" + IntegerToRomanBrute(number-90)
	}
	if number >= 50 {
		return "L" + IntegerToRomanBrute(number-50)
	}
	if number >= 40 {
		return "XL" + IntegerToRomanBrute(number-40)
	}
	if number >= 10 {
		return "X" + IntegerToRomanBrute(number-10)
	}
	if number == 9 {
		return "IX"
	}
	if number >= 5 {
		return "V" + IntegerToRomanBrute(number-5)
	}
	if number >= 4 {
		return "IV" + IntegerToRomanBrute(number-4)
	}
	return strings.Repeat("I", number)
}
