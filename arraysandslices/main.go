package arraysandslices

import "fmt"

func Sum(numbers [5]int) int {
	sum := 0
	for index, number := range numbers {
		fmt.Println("index", index, "number", number)
		sum += number
	}
	return sum
}
