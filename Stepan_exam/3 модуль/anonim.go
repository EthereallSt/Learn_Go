package main

import (
	"fmt"
	"log"
)

func main() {
	a := test(45)
	fmt.Println(a)
}
func test(a uint) uint {
	outputStr := 0
	for _, r := range string(a) {
		if r != 0 && r%2 != 0 {
			outputStr += int(r)
			log.Println(outputStr)
		}
	}
	return uint(outputStr)
}

//b, err := strconv.ParseUint(outputStr, 10, 32)
 