package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan<- string, wg *sync.WaitGroup) {
		ch <- "Sanju"
		ch <- "Sameer"
		ch <- "Karan"
		close(ch)
		wg.Done()
	}(myCh, wg)

	for val := range myCh {
		fmt.Println(val)
	}
	wg.Wait()
}
