package main

import "fmt"

var k, a, sum1 int

func main() {
	fmt.Scan(&k)
	for k > 0 {
		fmt.Scan(&a)
		k = k - 1
		if a > 9 && a < 100 && a%8 == 0 {
			sum1 = sum1 + a
		}
	}
	fmt.Println(sum1)
}
