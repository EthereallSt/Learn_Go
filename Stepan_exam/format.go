/*
На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу:
число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой.
Но если число равно или меньше нуля - выводить: "число R не подходит",
где R - исходное число с 2 цифрами после запятой и с 2 по ширине.
А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении .
*/

package main

import "fmt"

var a float32

func main() {
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Printf("число %2.2f не подходит", a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else if a > 0 && a <= 1000 {
		fmt.Printf("%.4f", a*a)
	}
}
