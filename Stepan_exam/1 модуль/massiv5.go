package main

import "fmt"

func main() {
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a
	}
	for i := 0; i < 3; i++ {
		var b, c int
		fmt.Scan(&b, &c)
		workArray[b], workArray[c] = workArray[c], workArray[b]
	}
	for _, value := range workArray {
		fmt.Print(value, " ")
	}
}

//input 99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
//output 99 137 151 187 117 71 20 187 93 67 type ok
