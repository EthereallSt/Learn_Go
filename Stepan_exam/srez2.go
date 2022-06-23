/*
На ввод подаются пять целых чисел, которые записываются в массив.
Найти и вывести максимальное число
*/

package main

import "fmt"

func main() {
	array := [5]int{}
	var a, max int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	fmt.Println(max)
	// ...
}
