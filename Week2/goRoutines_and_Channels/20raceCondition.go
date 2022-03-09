// go run --race 20raceCondition.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	statusCode := []int{}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		statusCode = append(statusCode, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		statusCode = append(statusCode, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		statusCode = append(statusCode, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(statusCode)
}
