/*
Идёт k-я секунда суток. Определите,
сколько целых часов h и целых минут m прошло с начала суток.
*/

package main

import "fmt"

var hours, min, a int

func main() {
	fmt.Scan(&a)
	hours = a / 3600
	min = (a - (hours * 3600)) / 60
	//sec = a - ((hours * 3600) + (min * 60))
	//fmt.Println(hours, min, sec)

	fmt.Printf("It is %d hours %d minutes.", hours, min)
}
