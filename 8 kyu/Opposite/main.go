package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(Opposite(number))
}

func Opposite(value int) int {
	return -value
}
