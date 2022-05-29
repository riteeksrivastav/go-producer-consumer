package main

import (
	"fmt"
	"sync"
)

func Producer(i int, in chan<- int, wg *sync.WaitGroup) {
	for j := 0; j < 10; j++ {
		in <- j * j
	}
	wg.Done()
}

func Consumer(in <-chan int, wg *sync.WaitGroup) {
	for data := range in {
		fmt.Println(data)
	}
	wg.Done()
}

func main() {
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	pipe := make(chan int)

	for i := 0; i < 10; i++ {
		wp.Add(1)
		go Producer(i, pipe, wp)
	}

	for i := 0; i < 10; i++ {
		go Consumer(pipe, wc)
	}

	wp.Wait()
	close(pipe)
	wc.Wait()
}
