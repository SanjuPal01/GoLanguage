package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2]) //string is nothing but collection of bytes and bytes is uint8 hence we are getting unicode values (string follow unicode8)
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	var s2 string = " new String"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	s1 := "Data to send"
	b := []byte(s1)
	// when we want to send our data then we convert it into bytes of data hence this function is very useful
	fmt.Printf("%v, %T\n", b, b)
}
