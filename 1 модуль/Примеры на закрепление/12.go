package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := 2

	if a < b && a != 0 {
		fmt.Println(1)
		return
	}
	if a == b {
		fmt.Println(1)
		return
	}
	if a == 0 {
		fmt.Println(a, " ")
	}
	for i := 0; i <= a; i++ {
		if i == 1 {
			fmt.Print(i, " ")
		}
		if i == b {
			fmt.Print(i, " ")
			b = b * 2
		}
	}
}
