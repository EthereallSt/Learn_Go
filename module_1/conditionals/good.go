//Определите является ли билет счастливым. Счастливым считается билет,
//в шестизначном номере которого сумма первых трёх цифр
//совпадает с суммой трёх последних.

package main

import "fmt"

func good() {
	var a, sum1, sum2 int
	fmt.Scan(&a)

	sum1 = a / 1000 // разделим число на 2 числа по 3 знака
	sum2 = a % 1000

	sum1 = sum1%10 + (sum1/10)%10 + (sum1 / 100) // получаем сумму 3 чисел
	sum2 = sum2%10 + (sum2/10)%10 + (sum2 / 100)

	//блок действий по сравнению чисел
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
