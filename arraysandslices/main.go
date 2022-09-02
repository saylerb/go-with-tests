package arraysandslices

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for index, number := range numbers {
		fmt.Println("index", index, "number", number)
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbersToSum := len(numbersToSum)

	sums := make([]int, lengthOfNumbersToSum)

	for index, slice := range numbersToSum {
		sums[index] += Sum(slice)
	}

	return sums
}
