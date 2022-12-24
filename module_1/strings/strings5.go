package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	var count, examp int
	fmt.Scan(&password)
	runes := []rune(password)
	lenth := len(runes)
	count = lenth
	for _, v := range runes {
		if lenth < 5 {
			fmt.Println("Wrong password")
			return
		}
		examp += 1
		if !(unicode.Is(unicode.Latin, v) || unicode.Is(unicode.Digit, v)) {
			count = -1
		}
	}
	fmt.Println(count, examp)
	if count == examp {
		fmt.Println("Ok")
	} else {package main

		import (
			"fmt"
		"unicode"
		)
		func main() {
			var password string
			var count, examp int
			fmt.Scan(&password)
			runes := []rune(password)
			lenth := len(runes)
			count = lenth
			for _, v := range runes {
				if lenth < 5 {
					fmt.Println("Wrong password")
					return
				}
				examp += 1
				if !(unicode.Is(unicode.Latin, v) || unicode.Is(unicode.Digit, v)) {
					count = -1
				}
			}
			if count == examp {
				fmt.Println("Ok")
			} else {
				fmt.Println("Wrong password")
			}
		}
		fmt.Println("Wrong password")
	}
}
