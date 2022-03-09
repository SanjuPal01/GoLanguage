package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hii Sanju Here"
	fmt.Println(str)
	pos := strings.Index(str, "Sanju")
	fmt.Println("Starting Index of Sanju is", pos)
	fmt.Println("Number of times H's appears", strings.Count(str, "H"))

	newStr := strings.Repeat(str, 3)
	fmt.Println(newStr)
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println("String contains 'Sanju'?  ", strings.Contains(str, "Sanju"))
}
