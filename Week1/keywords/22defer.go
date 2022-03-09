package main

import "fmt"

// you use defer in which function it will executes after that function execution
// we generally use defer for closing the resources
// defer works in LIFO format
func main() {
	fmt.Println("Sanju")
	defer fmt.Println("Sameer")
	fmt.Println("Karan")
	test()
	fmt.Println("End")

	a := "ABC"
	defer fmt.Println(a) //it will print the old value of a irrespective of changing it's value later
	a = "DEF"
}

func test() {
	defer fmt.Println("Heya")
}
