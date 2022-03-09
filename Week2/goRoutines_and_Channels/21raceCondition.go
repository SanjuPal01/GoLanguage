// go run --race 21raceCondition.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	statusCode := []int{}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		statusCode = append(statusCode, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		statusCode = append(statusCode, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		statusCode = append(statusCode, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()

	mut.RLock()
	fmt.Println(statusCode)
	mut.RUnlock()
}
