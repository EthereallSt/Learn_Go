package main

import (
	"fmt"
	"time"
)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go printer(outputStream)

	for _, i := range "112334456" {
		inputStream <- string(i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	lastAnswer := "1488"
	answer := <-inputStream
	if answer != lastAnswer {
		outputStream <- answer
	}
	lastAnswer = answer
	close(outputStream)
}
