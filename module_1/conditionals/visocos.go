/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100
*/

package main

import "fmt"

func visocos() {
	var god int

	fmt.Scan(&god) // считываем год

	if god%400 == 0 || god%4 == 0 && god%100 != 0 { // сравниваем
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
