package main

import "fmt"

var a, b int

func main() {
	srez := []int{} //объявляю срез
	fmt.Scan(&b)
	for b > 0 {
		fmt.Scan(&a)
		srez = append(srez, a) //добавляю каждый раз к срезу новый элемент а
		b--
	}
	fmt.Println(srez[3]) //вывожу 4ый элемент среза
}
