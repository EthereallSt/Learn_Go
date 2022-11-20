package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s string

func main() {
	//s = "1 149,6088607594936;1 179,0666666666666"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	digit := rune(';')
	digit1 := []rune{}
	digit2 := []rune{}
	tumbler := 0
	s2 := strings.Replace(s, ",", ".", -1)

	for _, r := range s2 { // прохожу строку и разбиваю на 2 числа по рунам
		if r == digit {
			tumbler = 1
			continue
		}
		if tumbler == 0 && r != ' ' {
			digit1 = append(digit1, r)
		}
		if tumbler == 1 && r != ' ' {
			digit2 = append(digit2, r)
		}
	}
	digit3 := string(digit1) //Преобразую срез рун в строку для дальнейших преобразований
	digit4 := string(digit2)

	floatdigit3, err := strconv.ParseFloat(digit3, 32) //перевожу во флоат
	if err != nil {
		fmt.Println("ошибка в 1")
	}
	floatdigit4, err := strconv.ParseFloat(digit4, 32)
	if err != nil {
		fmt.Println("ошибка в 2")
	}
	result := floatdigit3 / floatdigit4 //получаю частное
	fmt.Printf("%.4f", result)          //делаю вывод с 4 знаками после запятой
}
