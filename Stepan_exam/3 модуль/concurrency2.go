package main

func main() {

}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(ch chan int) {
		defer close(ch)

		select {
		case answer := <-firstChan:
			ch <- answer * answer
		case answer := <-secondChan:
			ch <- answer * 3
		case <-stopChan:
		}
	}(output)
	return output
}
