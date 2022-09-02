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
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			withoutHead := numbers[1:]
			result = append(result, Sum(withoutHead))
		}
	}
	return result
}
