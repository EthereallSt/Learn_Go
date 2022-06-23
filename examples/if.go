/*
// использование \n позволяет сделать перенос строки
var a1 string = "123"
var a2 string = "1234"
fmt.Printf("%q \n%s", a1, a2)
// вывод:
// "123"
// 1234
*/

package main

import "fmt"

func main() {
	var a int // Вводим  переменную а
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
