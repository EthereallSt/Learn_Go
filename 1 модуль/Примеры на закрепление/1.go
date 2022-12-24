package main

import "fmt"

var a, summ, prom int

func main() {
	fmt.Scan(&a)
	for a > 10 {
		prom = a % 10
		summ = summ + prom
		a = a / 10
	}
	summ = summ + a
	fmt.Println(summ)
}
