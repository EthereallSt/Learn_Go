package main

import "fmt"

var text string

func main() {
	fmt.Scan(&text)
	f(text)
}
func f(text string) string {
	fmt.Println(text)
	return ""
}
