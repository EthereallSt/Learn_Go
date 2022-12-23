package main

import "fmt"

var a int

func main() {
	for fmt.Scan(&a); a <= 100; fmt.Scan(&a) {
		if a < 10 {
			continue
		}
		fmt.Println(a)
	}
}
