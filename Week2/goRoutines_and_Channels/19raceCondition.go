// go run --race 19raceCondition.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	statusCode := []int{}
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		statusCode = append(statusCode, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		statusCode = append(statusCode, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		statusCode = append(statusCode, 3)
		wg.Done()
	}(wg)
	wg.Wait()
	fmt.Println(statusCode)
}
