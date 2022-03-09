package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Sanju")
	greet("Sahiba")
}

func greet(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(i, str)
	}
}
