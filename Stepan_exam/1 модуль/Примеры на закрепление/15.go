package main

import (
	"fmt"
	"strconv"
)

var number int

func main() {
	fmt.Scan(&number)
	numberString := strconv.Itoa(number)
	for _, v := range numberString {
		a, _ := strconv.Atoi(string(v))
		fmt.Print(a * a)
	}
}
