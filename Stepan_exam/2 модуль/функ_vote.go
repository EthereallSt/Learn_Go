package main

func main() {
	vote(0, 1, 1)
}

func sumInt(args ...int) (sum, kol int) { //Функция принимающая на вход любое кол-во аргументов.
	for _, value := range args { //Она же считает кол-во Запомнить for _, value:= range args
		sum += value
		kol++
	}
	return kol, sum
}

func vote(x int, y int, z int) int {
	if x == y || x == z {
		return x
	} else {
		return z
	}
}

func fibbonachi(n int) int {
	for i; i < n; i := 0 {
		n = n + nn
		i++
	}
}
