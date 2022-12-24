package main

import "fmt"

var n, digit int

func main() {
	for i := 30; i > 0; i-- {
		fmt.Scan(&n)
		if n > 5 && n < 21 {
			fmt.Printf("%d korov", n)
			return
		}
		digit = n % 10
		switch {
		case digit == 1:
			fmt.Printf("%d korova", n)
		case digit > 4:
			fmt.Printf("%d korov", n)
		case digit == 2 || digit == 3 || digit == 4:
			fmt.Printf("%d korovy", n)
		default:
			fmt.Printf("%d korov", n)
		}
	}
}
