package main

import (
	"fmt"
	"strings"
)

func main() {
	var word, podword string
	fmt.Scan(&word, &podword)

	index := strings.Index(word, podword)
	fmt.Println(index)
}
