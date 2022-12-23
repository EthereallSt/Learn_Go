package main

import "fmt"

var a, b, c int

func main() {
	fmt.Scan(&a, &b, &c)
	if a+b > c && a+c > b && b+c > a {
		fmt.Print("Существует")
	} else {
		fmt.Print("Не существует")
	}
}
