package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wwg := new(sync.WaitGroup)

	for i := 0; i < 6; i++ {
		wwg.Add(1)
		go work(i, wwg)
	}

	wwg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("ВСе горутины завершили выполнение")
}

func work(id int, wwg *sync.WaitGroup) {
	id = id + 2
	defer wwg.Done()
	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}
