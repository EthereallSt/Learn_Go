package main

import "fmt"

func main() {
	args := make(chan int)
	stop := make(chan struct{})

	result := calculator(args, stop)
	args <- 3
	args <- 5
	args <- 2
	args <- 2
	args <- 2
	args <- 8
	// stop <- struct{}{}
	close(stop)

	fmt.Println(<-result)

}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	answer := make(chan int)
	go func(answer chan int) {
		defer close(answer)
		var sum int
		for {
			select {
			case answer := <-arguments:
				sum += answer
			case <-done:
				answer <- sum
			}
		}
	}(answer)
	return answer
}
