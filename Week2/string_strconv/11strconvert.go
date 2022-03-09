package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// str := "Hey Sanju here 34"
	str := "34"
	fmt.Println("The size of Integer is:", strconv.IntSize)
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}
