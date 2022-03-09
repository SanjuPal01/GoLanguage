package main

import "fmt"

func main() {
	n := 3.14
	fmt.Printf("%v, %T\n", n, n)
	var m float32 = 5.4
	fmt.Printf("%v, %T\n", m, m)

	a := 13.7e72
	b := 12.3e45
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	p := 10.4
	q := 5.4

	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)

	d := 3.45678
	fmt.Printf("%.2f\n", d)
}
