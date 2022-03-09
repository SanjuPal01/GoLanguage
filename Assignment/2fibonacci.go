package main

import "fmt"

func main() {
	fmt.Print("Enter Number: ")
	var num int
	fmt.Scanln(&num)
	generateFibonacci(num)
}
func generateFibonacci(num int) {
	first := 0
	second := 1
	for i := 0; i < num; i++ {
		fmt.Print(first, " ")
		third := first + second
		first = second
		second = third
	}
	fmt.Println()
}
