package main

import "fmt"

var a, b, max int

func main() {
	if b == a {
		if b%7 == 0 {
			max = b
		}
	}
	fmt.Scan(&a, &b)
	for b >= a {
		if b%7 == 0 {
			max = b
			fmt.Println(max)
			break
		}
		b--
	}
	if max != b {
		fmt.Println("NO")
	}

}
