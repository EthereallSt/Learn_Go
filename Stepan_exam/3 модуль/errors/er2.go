package main

import (
	"fmt"
)

var s string

func main() {
	fmt.Scan(&s)
	length := len(s)*2 - 1
	buffer := make([]byte, 0, length)

	for i := range s {
		buffer = append(buffer, s[i])

		if len(buffer) < length {
			buffer = append(buffer, '*')
		}

	}
	fmt.Println(string(buffer))
}
