package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	go greet("Sanju", wg)
	wg.Add(1)
	greet("Sahiba", nil)
	wg.Wait()
}

func greet(str string, wg *sync.WaitGroup) {
	for i := 0; i < 6; i++ {
		fmt.Println(i, str)
	}
	if wg != nil {
		defer wg.Done()
	}
}
