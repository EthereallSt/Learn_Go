package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scan(&word)
	norm := 1

	for _, v := range word {
		if norm == strings.Count(word, string(v)) {
			fmt.Print(string(v))
		} else {
			continue
		}

	}
}
