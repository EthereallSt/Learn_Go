package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var answer int

	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	scanner := bufio.NewScanner(os.Stdin) // cоздал сканер
	for scanner.Scan() {
		s := scanner.Text()
		chislo, _ := strconv.Atoi(s)
		answer += chislo
	}

	io.WriteString(os.Stdout, strconv.Itoa(answer))
}
