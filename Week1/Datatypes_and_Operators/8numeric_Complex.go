package main

import "fmt"

func main() {
	a := 1 + 2i
	fmt.Printf("%v, %T\n", a, a)

	b := 2i
	fmt.Printf("%v, %T\n", b, b)

	var c complex64 = 3 + 4i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	d := complex(5, 6)
	fmt.Printf("%v, %T\n", d, d)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

}
