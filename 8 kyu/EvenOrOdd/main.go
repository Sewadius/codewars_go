package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(EvenOrOdd(number))
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
