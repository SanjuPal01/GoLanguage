package main

import (
	"fmt"
)

const p = -1
const q = iota
const r = iota

const (
	s = iota
	t
	u
)
const (
	g = iota
	h
	i
)

func main() {
	const myConst int = 42
	// const myNewConst float64 = math.Sin(1.57)			- Will give error because function call can happen only at runtime
	// and constant value is replaced at their places in compile time
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 4
	const b string = "foo"
	const c float32 = 4.65
	const d bool = true

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	const f = 42
	fmt.Printf("%v, %T\n", f, f)

	fmt.Println(p, q, r)
	fmt.Println(s, t, u)
	fmt.Println(g, h, i)
}
