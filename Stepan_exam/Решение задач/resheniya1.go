package main

import "fmt"

var s string

func main() {
	fmt.Scan(&s)
	rs := []rune(s)

	for i := range rs {
		if rs[i]%2 == 0 {
			rs[i] = '*'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))
}
