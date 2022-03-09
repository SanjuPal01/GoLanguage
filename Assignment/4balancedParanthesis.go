package main

import (
	"fmt"

	"example.com/stack"
)

func main() {
	fmt.Print("Enter Paranthesis for checking it's Balanced or Not: ")
	var str string
	fmt.Scanln(&str)
	if isBalanced(str) {
		fmt.Println(str, "is Balanced")
	} else {
		fmt.Println(str, "is Not Balanced")
	}
}

func isMatching(a int, b int) bool {
	if a == '(' && b == ')' {
		return true
	} else if a == '[' && b == ']' {
		return true
	} else if a == '{' && b == '}' {
		return true
	} else {
		return false
	}
}

func isBalanced(str string) bool {
	st := stack.Init(10)
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			st.Push(int(str[i]))
		} else {
			if st.IsEmpty() {
				return false
			} else if isMatching(st.Peek(), int(str[i])) == false {
				return false
			} else {
				st.Pop()
			}
		}
	}
	return st.IsEmpty() == true
}
