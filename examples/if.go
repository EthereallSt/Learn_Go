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
