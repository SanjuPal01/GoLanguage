package main

import "fmt"

var i int = 5 //first letter lowerCase - for package scope
var J int = 6 //first letter UpperCase - to export

// Defining too much variable outside the main function would affect on its readibility
// So what we can do we can create one var block so we can avoid this var keyword(Everytime when we are declaring and defining the variables)

var (
	a int = 1
	b int = 2
	c int = 3
)

func main() {
	fmt.Println(i)
	var i int = 7
	fmt.Println(i)
	fmt.Println(J)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
