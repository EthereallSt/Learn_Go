package main

import "fmt"

var n, dig_c int

func main() {
	fmt.Scan(&n)
	dig_c = (n-1)%9 + 1
	fmt.Println(dig_c)
}
