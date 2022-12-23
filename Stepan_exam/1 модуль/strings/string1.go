package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var a string

var last = "."

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	rs := []rune(text)
	if unicode.IsUpper(rs[0]) == true {
		strings.TrimSuffix(text, "\n")
		if strings.HasSuffix(text, last) == true {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	} else {
		fmt.Println("Wrong")
	}

}
