package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 5.6
	fmt.Println(i)
	j := int(i)
	fmt.Println(j)

	k := 5
	fmt.Println(k)
	l := float32(k)
	fmt.Println(l)

	// WRONG WAY
	var val int = 65
	val2 := string(val)
	fmt.Println(val2)

	// For doing conversions in string we use "strconv" package
	var newVal int = 66
	newVal2 := strconv.Itoa(newVal)
	fmt.Println(newVal2)
}
