package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hiii")
	// }

	// while loop
	i := 0
	for i <= 4 {
		fmt.Println(i)
		i++
	}

	arr := []string{"Sanju", "Pal"}
	for i, j := range arr {
		fmt.Println(i, j)
	}

	for i, j := range "Sanju" {
		fmt.Printf("%d, %c\n", i, j)
	}

	// Iterating Maps
	myMap := map[int]string{
		0: "Sanju",
		1: "Sameer",
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
