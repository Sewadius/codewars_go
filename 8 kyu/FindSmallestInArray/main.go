package main

import "fmt"

func main() {
	var numbers = []int{34, -345, -1, 100}
	fmt.Println(SmallestIntegerFinder(numbers))
}

func SmallestIntegerFinder(numbers []int) int {
	minimum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < minimum {
			minimum = numbers[i]
		}
	}
	return minimum
}
