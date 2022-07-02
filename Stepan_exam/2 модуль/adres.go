package main		// Меняет значения адреса переменных местами 

import "fmt"

var x1, x2, x3 *int

func main() {
	fmt.Print("")
	test(x1, x2)
}

func test(x1 *int, x2 *int) {
	var x3 *int
	x3 = x1
	x1 = x2
	x2 = x3
	fmt.Println(x1, x2)
}
