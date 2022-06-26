// НЕ закончил
package main

import "fmt"

var n, digit int

func main() {
	fmt.Scan(&n)
	digit = n % 10
	switch {
	case digit == 1:
		fmt.Printf("%d korova", n)
	case 4 < n:
		if n < 21 {
			fmt.Printf("%d korov", n)
		}
	case digit == 2 || digit == 3 || digit == 4:
		fmt.Printf("%d korovy", n)
	default:
		fmt.Printf("%d korov", n)
	}
}
