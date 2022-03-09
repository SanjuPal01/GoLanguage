package main

import "fmt"

// You can't do pointer arithmetic in golang
func main() {
	var a int = 42
	b := a //Pass by value
	fmt.Println(a, b)
	b = 25
	fmt.Println(a, b)

	// Pointers
	c := 100
	// 1st way - to define pointers
	var d *int = &c
	// d = &c
	// 2nd way
	e := &c
	fmt.Println(d, e)
	fmt.Println(*d, *e)
	*e = 25
	fmt.Println(c, *d, *e)

	// Arrays
	arr := [3]int{1, 2, 3}
	val1 := &arr[0]
	val2 := &arr[1]
	fmt.Printf("%v %p %p\n", arr, val1, val2)
	fmt.Printf("%v %v %v\n", arr, *val1, *val2)
}
