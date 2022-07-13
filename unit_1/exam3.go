/*https://stepik.org/lesson/228261/step/16?unit=200794*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println("It is", num/30, "hours", 2*(num%30), "minutes.")
}
