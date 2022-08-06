/*https://stepik.org/lesson/228263/step/8?unit=200796*/

package main

import "fmt"

func main() {
	var num int
	for fmt.Scan(&num); num <= 100; fmt.Scan(&num) {
		if num > 9 && num < 101 {
			fmt.Println(num)
		}
	}
}
