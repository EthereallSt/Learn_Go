package main

import "fmt"

var a, b, sr_chislo float32

func main() {
	fmt.Scan(&a, &b)
	sr_chislo = (a + b) / 2
	fmt.Print(sr_chislo)
}
