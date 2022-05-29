package main

import (
	"fmt"
	"sync"
)

func Producer(in chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	wg.Done()
}

func Consumer(in <-chan int, wg *sync.WaitGroup) {
	for data := range in {
		fmt.Println(data)
	}
	wg.Done()
}

//This implementation has a problem, the producer is closing the chanenl, if we extend the one producer to
// multi producer we will have to close the channel out side of producer and that can't be solved by single wg.
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	pipe := make(chan int)
	go Producer(pipe, wg)
	go Consumer(pipe, wg)
	wg.Wait()
}
