/*https://stepik.org/lesson/228261/step/15?unit=200794*/

package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	fmt.Println(string(num[len(num)-2]))
}
