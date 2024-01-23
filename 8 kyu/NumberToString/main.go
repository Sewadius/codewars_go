package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(NumberToString(number))
}

func NumberToString(n int) string {
	return strconv.FormatInt(int64(n), 10)
}
