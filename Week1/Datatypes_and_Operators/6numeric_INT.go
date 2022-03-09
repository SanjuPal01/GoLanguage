package main

import (
	"fmt"
)

func main() {
	var i int = 4
	var j uint = 4
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var p int = 4
	var q int32 = 5
	//fmt.Println(p + q) //Wrong (They need to have same type)
	fmt.Println(p + int(q))

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b) //if none of bit set then put 1 else 0 - you can say negation of OR

	a = 8
	fmt.Println(a << 3)
	fmt.Println(a >> 3)

}
