package main

import (
	"fmt"
)

const pi = 3.14

func maini() {
	printArea()
}

func printArea() {
	raduis := 2
	area := (float32(raduis) * float32(raduis) * pi)

	fmt.Printf("Радиус круга = %d см\n", raduis)
	fmt.Println("Площадь круга = радус в квадрате * на число pi")
	fmt.Printf("Площадь круга =%.2f см. кв\n", area)
}
