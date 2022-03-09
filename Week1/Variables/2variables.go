package main

import "fmt"

func main() {
	// 1st way to define the variables
	var i int
	fmt.Println(i) //By default value is 0
	i = 5
	fmt.Println(i)

	// 2nd way
	var j float32 = 6
	fmt.Println(j)

	// 3rd way - Automatically detects the type
	k := 7
	fmt.Printf("%v, %T\n", k, k) //%v - Value , %T - Type
}
