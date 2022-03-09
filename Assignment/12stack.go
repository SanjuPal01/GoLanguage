package main

import (
	"fmt"

	"example.com/stack"
)

func main() {
	st := stack.Init(10)
	fmt.Println(st.Size())
	st.Push(12)
	fmt.Println(st.Size())
	fmt.Println(st.Peek())
	fmt.Println(st.Pop())
	fmt.Println(st.Size())

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	fmt.Println(st.Size())
	fmt.Println(st.Peek())
	fmt.Println(st.Pop())
	fmt.Println(st.Size())

}
