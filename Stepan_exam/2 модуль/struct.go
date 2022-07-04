package main



var c Circle

type Circle struct { // объявление (инициализация) структуры Circle
	x, y, r float64
}

func m() Circle {
	c := Circle{}
	return c
}

func main() {
	fmt.Println(c)
}
