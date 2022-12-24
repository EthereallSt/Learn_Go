/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4..).
*/

package main

import (
	"fmt"
)

var a, n int

func main() {
	fmt.Scan(&n)
	array := []int{}
	for i := 0; n > 0; {
		fmt.Scan(&a)
		array = append(array, a)
		if i%2 == 0 {
			fmt.Print(" ", array[i])
		}
		i++
		n--
	}
}
