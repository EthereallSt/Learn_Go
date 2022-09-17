package main

import (
	"fmt"
)

var s string
var max = 0

func main() {
	fmt.Scan(&s)
	for i, c := range s {
		if i >= max {
			max = byte(i)
		}
		fmt.Println(max, с)
	}

}
