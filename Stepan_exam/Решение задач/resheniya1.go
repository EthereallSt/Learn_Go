package main

import "fmt"

var s string
var c int

func main() {
	fmt.Scan(&s)
	rs := []byte(s)

	for c, i := range rs {
		if c%2 == 0 {
			rs[i] = '*'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))
}
