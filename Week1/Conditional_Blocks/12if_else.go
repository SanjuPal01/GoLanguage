package main

import "fmt"

func main() {
	fmt.Print("Enter Number: ")
	var num int
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
