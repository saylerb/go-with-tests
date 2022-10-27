package main

import "testing"

func TestRoman(t *testing.T) {
	AssertRoman(IntegerToRoman(0), "", t)
	AssertRoman(IntegerToRoman(1), "I", t)
	AssertRoman(IntegerToRoman(2), "II", t)
	AssertRoman(IntegerToRoman(3), "III", t)
	AssertRoman(IntegerToRoman(4), "IV", t)
	AssertRoman(IntegerToRoman(5), "V", t)
	AssertRoman(IntegerToRoman(6), "VI", t)
	AssertRoman(IntegerToRoman(7), "VII", t)
	AssertRoman(IntegerToRoman(8), "VIII", t)
	AssertRoman(IntegerToRoman(9), "IX", t)
	AssertRoman(IntegerToRoman(10), "X", t)
	AssertRoman(IntegerToRoman(11), "XI", t)
	AssertRoman(IntegerToRoman(12), "XII", t)
	AssertRoman(IntegerToRoman(13), "XIII", t)
	AssertRoman(IntegerToRoman(14), "XIV", t)
	AssertRoman(IntegerToRoman(15), "XV", t)
	AssertRoman(IntegerToRoman(16), "XVI", t)
	AssertRoman(IntegerToRoman(17), "XVII", t)
	AssertRoman(IntegerToRoman(18), "XVIII", t)
	AssertRoman(IntegerToRoman(19), "XIX", t)
	AssertRoman(IntegerToRoman(20), "XX", t)
	AssertRoman(IntegerToRoman(40), "XL", t)
	AssertRoman(IntegerToRoman(45), "XLV", t)
	AssertRoman(IntegerToRoman(50), "L", t)
	AssertRoman(IntegerToRoman(51), "LI", t)
	AssertRoman(IntegerToRoman(90), "XC", t)
	AssertRoman(IntegerToRoman(93), "XCIII", t)
	AssertRoman(IntegerToRoman(400), "CD", t)
	AssertRoman(IntegerToRoman(410), "CDX", t)
	AssertRoman(IntegerToRoman(500), "D", t)
	AssertRoman(IntegerToRoman(900), "CM", t)
	AssertRoman(IntegerToRoman(1000), "M", t)
	AssertRoman(IntegerToRoman(1001), "MI", t)
	AssertRoman(IntegerToRoman(1990), "MCMXC", t)
	AssertRoman(IntegerToRoman(2007), "MMVII", t)
	AssertRoman(IntegerToRoman(2008), "MMVIII", t)
}

func AssertRoman(got string, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
