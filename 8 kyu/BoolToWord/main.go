package main

import "fmt"

func main() {
	var boolean bool
	fmt.Scan(&boolean)
	fmt.Println(BoolToWord(boolean))
}

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}
