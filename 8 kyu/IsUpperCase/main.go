package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type MyString string

func main() {
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()

	fmt.Println(MyString(text).IsUpperCase())
}

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if !unicode.IsUpper(char) && unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
