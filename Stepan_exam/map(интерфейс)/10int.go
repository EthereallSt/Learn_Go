package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(map[int]int)  //cоздаю и инициализирую мапу(интерфейс)
	keys := make([]int, 0, 10) // создаю список из 10 цифр 0, в него сложу все поступившие цифры
	digit := 0                 // переменная для поступающих цифр
	amount := 0                // переменная со значением - 10 (для кол-ва)

	for fmt.Scan(&digit); amount < 10; fmt.Scan(&digit) { // В это цикле я сканирую цифры
		amount++
		keys = append(keys, digit) // добавляю их в список кейс
	}
	for _, key := range keys { // В этом цикле я прохожу по списку кейс.
		if _, inMap := nums[key]; !inMap { // Чтобы внутри него пройти список намс и при отсутствии
			nums[key] = work(key) // клю1а, который совпадает со значением в кейс,
		} // добавить в намс результат работы функции ворк
		fmt.Print(nums[key], " ")
	}
}
func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}
