package main

import "fmt"

var a, b, i int

func main() {
	i = 0
	fmt.Scan(&b)
	for b > 0 {
		fmt.Scan(&a)
		if a == 0 {
			i++
		}
		b--
	}
	fmt.Println(i)
}
