package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Multiply(a, b))
}

func Multiply(a, b int) int {
	return a * b
}
