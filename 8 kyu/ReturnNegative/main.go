package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(MakeNegative(x))
}

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}
