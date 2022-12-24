package main

import (
	"fmt"
	"math"
)

var a, b float64

func main() {
	var a, b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	a = float64(a)
	b = float64(b)

	fmt.Println(math.Hypot(a, b))
}
