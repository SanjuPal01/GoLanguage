package main

import (
	"fmt"

	"example.com/stack"
)

func main() {
	fmt.Print("Enter Valid Infix Statement: ")
	var infix string
	fmt.Scanln(&infix)
	fmt.Println("Infix: ", infix)
	postfix := infixtoPostfix(infix)
	fmt.Println("Postfix: ", postfix)
}

func precedence(c int) int {
	if c == '^' {
		return 3
	} else if c == '*' || c == '/' {
		return 2
	} else if c == '+' || c == '-' {
		return 1
	} else {
		return -1
	}
}

func infixtoPostfix(str string) string {
	st := stack.Init(100)
	var res string
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' || str[i] >= '0' && str[i] <= '9' {
			res += string(str[i])
		} else if str[i] == '(' {
			st.Push(int(str[i]))
		} else if str[i] == ')' {
			for !st.IsEmpty() && st.Peek() != '(' {
				res += string(st.Pop())
			}
			st.Pop()
		} else {
			for !st.IsEmpty() && precedence(int(str[i])) <= precedence(st.Peek()) {
				res += string(st.Pop())
			}
			st.Push(int(str[i]))
		}
	}
	for !st.IsEmpty() {
		res += string(st.Pop())
	}
	return res
}
