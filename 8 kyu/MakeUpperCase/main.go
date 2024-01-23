package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(MakeUpperCase(input))
}

func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
