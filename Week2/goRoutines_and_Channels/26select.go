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

	//Or you can use infinite for loop

	for x := 0; x < 2; x++ {
		select {
		case rec1 := <-one:
			fmt.Println(rec1)
		case rec2 := <-two:
			fmt.Println(rec2)
		}
	}

}
