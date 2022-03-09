package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	Panicker()
	fmt.Println("End")
}
func Panicker() {
	fmt.Println("About to Panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err)
		}
	}()
	panic("Bad Happened")
	fmt.Println("Done Panicking")
}
