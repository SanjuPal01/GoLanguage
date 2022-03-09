package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "hello"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		two <- "hii"
	}()

	select {
	case rec1 := <-one:
		fmt.Println(rec1)
	case rec2 := <-two:
		fmt.Println(rec2)
		// default:
		// 	fmt.Println("Not Received Yet")
	}
}
