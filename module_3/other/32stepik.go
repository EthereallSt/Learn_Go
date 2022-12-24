package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	answer := adding("1t3", "0KJDFhfbefhkj9")
	fmt.Println(answer)
}
func adding(first string, second string) int64 {
	c := ""
	for _, r := range first {
		if unicode.IsDigit(r) == true {
			c = c + string(r)
		}
	}
	log.Println(c)
	d := ""
	for _, r := range second {
		if unicode.IsDigit(r) == true {
			d = d + string(r)
		}
	}
	log.Println(d) //Теперь вместо first & second у нас c & d

	a, err := strconv.Atoi(c)
	if err != nil {
		log.Println("PZDC")
	}
	b, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println(err)
	}
	return int64(a + b)
}
