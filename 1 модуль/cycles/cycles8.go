package main

import (
	"fmt"
	"strconv"
)

var a, b int = 0, 0

func main() {
	fmt.Scan(&a, &b)

	as := strconv.Itoa(a)
	bs := strconv.Itoa(b)

	for _, i := range as {
		for _, is := range bs {
			if i == is {
				fmt.Printf("%v ", string(i))
			}
		}
	}
}
