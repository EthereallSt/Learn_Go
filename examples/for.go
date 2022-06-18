/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/

package main

import "fmt"

var pos, maxpos, i int

func main() {
	i := 1
	fmt.Scan(&pos)
	for pos != 0 {
		if pos > maxpos {
			maxpos = pos
			i = 1
		} else if pos == maxpos {
			i = i + 1
		}
		fmt.Scan(&pos)
	}
	fmt.Println(i)
}
