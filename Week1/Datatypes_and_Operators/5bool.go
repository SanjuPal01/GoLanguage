package main

import "fmt"

func main() {
	m := 1 == 1
	n := 1 == 2
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)

	var i bool
	fmt.Printf("%v, %T\n", i, i)
}
