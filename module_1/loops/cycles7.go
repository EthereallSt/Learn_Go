package main

import "fmt"

func main() {
	var money, procents, finisnMoney, year int = 0, 0, 0, 0
	fmt.Scan(&money, &procents, &finisnMoney)

	for money <= finisnMoney {
		money = money + money*procents/100
		year += 1
	}
	fmt.Println(year)
}
