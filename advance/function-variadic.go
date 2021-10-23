package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 24, 53, 44))

	numbers := []int{10, 24, 53, 44}
	total := sumAll(numbers...)
	fmt.Println(total)
}
