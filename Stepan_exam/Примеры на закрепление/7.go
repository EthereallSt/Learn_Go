/*
Найдите количество минимальных элементов в последовательности.
*/

package main

import "fmt"

var n, kol_vo, min, a int

func main() {
	min = 10000
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&a)
		if a < min {
			min = a
			kol_vo = 1
		} else if a == min {
			kol_vo++
		}
		n--
	}
	fmt.Println(kol_vo)
}
