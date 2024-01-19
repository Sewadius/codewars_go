package main

import "fmt"

func main() {
	var numbers []int = make([]int, 0, 5)
	numbers = append(numbers, 1, 2, 3, -6, -4)
	fmt.Println(PositiveSum(numbers))
}

func PositiveSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum += numbers[i]
		}
	}
	return sum
}
