package main

import "fmt"

func main() {
	fmt.Print("Enter Number: ")
	var num int
	fmt.Scanln(&num)
	/*
		if num == 1 {
			fmt.Println("One")
		} else if num == 2 {
			fmt.Println("Two")
		} else {
			fmt.Println("Neither One nor Two")
		}
	*/

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Neither One nor Two")
	}
}
