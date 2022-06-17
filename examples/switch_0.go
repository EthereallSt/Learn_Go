package main

import "fmt"

func switchi() {
	var a, bost_del, cost_del int
	fmt.Scan(&a)

	bost_del = a % 10        // последняя цифра
	cost_del = (a / 10) % 10 //предпоследняя цифра
	a = a / 100

	switch {

	case a != bost_del && a != cost_del && bost_del != cost_del:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
