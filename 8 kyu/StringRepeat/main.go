package main

import "fmt"

func main() {
	var s string
	var repeat int
	fmt.Scan(&s, &repeat)
	fmt.Println(RepeatStr(repeat, s))
}

func RepeatStr(repetitions int, value string) string {
	result := ""
	for i := 0; i < repetitions; i++ {
		result += value
	}
	return result
}
