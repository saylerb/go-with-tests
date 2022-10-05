package main

import (
	"fmt"
	"testing"
)

func rgbSimple(r int, g int, b int) string {
	// %02 padd with 2 zero's
	// %X uppercase hex representation
	return fmt.Sprintf("%02X%02X%02X", normalize(r), normalize(g), normalize(b))
}

func getRemainders(num int) (int, int) {
	result := num / 16
	remainder := num % 16
	remainderTwo := result % 16
	return remainder, remainderTwo
}

func toHexDigit(num int) string {
	if num > 10 {
		return string(rune(num - 10 + 65))
	} else {
		return fmt.Sprintf("%d", num)
	}
}

func toHexString(num int) string {
	remainder, remainderTwo := getRemainders(num)

	right := toHexDigit(remainder)
	left := toHexDigit(remainderTwo)

	return left + right
}

func rgb(r int, g int, b int) string {
	return toHexString(normalize(r)) +
		toHexString(normalize(g)) +
		toHexString(normalize(b))
}

func normalize(num int) int {
	if num > 255 {
		return 255
	} else if num < 0 {
		return 0
	} else {
		return num
	}
}

func TestRgbToHex(t *testing.T) {
	assertRgb(rgb(255, 255, 255), "FFFFFF", t)
	assertRgb(rgb(255, 255, 300), "FFFFFF", t)
	assertRgb(rgb(0, 0, 0), "000000", t)
	assertRgb(rgb(148, 0, 211), "9400D3", t)
}

func assertRgb(got, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
