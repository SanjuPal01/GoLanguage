package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 4) // Buffered channel - it will not give error even if won't take every data out from channel
	// and it will start ignoring the data once it reached it's storage
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Receive only : receive Data from channel only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// ch <- 3
		// close(ch)
		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("Channel Closed")
		}
		wg.Done()
	}(myCh, wg)

	// Send Only : Send data to channel only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// fmt.Println(<-ch)
		ch <- 2
		ch <- 3
		fmt.Println("Length: ", len(myCh))
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
