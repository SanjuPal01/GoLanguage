package main

import (
	"fmt"
	"sync"
)

/*
In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine.
*/
func main() {
	/*
		It will give error because we are not using channels in go routines
			// var myCh chan int
			myCh1 := make(chan int) // Unbuffered channel
			// myCh2 := make(chan int, 2)		//Buffered
			myCh1 <- 2
			myCh1 <- 3
			fmt.Println(<-myCh1)
			fmt.Println(<-myCh1)

	*/

	myCh := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 1
		wg.Done()

	}(myCh, wg)
	wg.Wait()
}
