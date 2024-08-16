package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := sumEverythingExcept(arr, 1)
	fmt.Println(sum)
}

func sumEverythingExcept(numbers []int, exceptNumber int) int {
	sum := 0

	for _, number := range numbers {
		if number != exceptNumber {
			sum += number
		}
	}

	return sum
}
