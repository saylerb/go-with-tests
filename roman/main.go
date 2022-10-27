package main

import "strings"

func IntegerToRoman(number int) string {
	if number >= 1000 {
		return "M" + IntegerToRoman(number-1000)
	}
	if number >= 900 {
		return "CM" + IntegerToRoman(number-900)
	}
	if number >= 500 {
		return "D" + IntegerToRoman(number-500)
	}
	if number >= 400 {
		return "CD" + IntegerToRoman(number-400)
	}
	if number >= 90 {
		return "XC" + IntegerToRoman(number-90)
	}
	if number >= 50 {
		return "L" + IntegerToRoman(number-50)
	}
	if number >= 40 {
		return "XL" + IntegerToRoman(number-40)
	}
	if number >= 10 {
		return "X" + IntegerToRoman(number-10)
	}
	if number == 9 {
		return "IX"
	}
	if number >= 5 {
		return "V" + IntegerToRoman(number-5)
	}
	if number == 4 {
		return "IV"
	}
	return strings.Repeat("I", number)
}
