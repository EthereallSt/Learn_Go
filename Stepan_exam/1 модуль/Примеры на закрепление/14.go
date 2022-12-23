package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number, digit int
	fmt.Scan(&number, &digit)

	numberString := strconv.Itoa(number)
	digitString := strconv.Itoa(digit)

	for _, ix := range numberString {
		if string(ix) == digitString {
			continue
		}
		fmt.Print(string(ix))
	}
}
