package main

import "fmt"

var one, two, three, a int

func main() {

	fmt.Scan(&a)
	one = a / 100
	three = a % 10
	two = (a / 10) % 10

	fmt.Printf("%d%d%d", three, two, one)
}
