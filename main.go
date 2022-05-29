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

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	pipe := make(chan int)
	go Producer(pipe, wg)
	go Consumer(pipe, wg)
	wg.Wait()
}
