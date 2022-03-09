package stack

import "fmt"

type Stack struct {
	arr []int
	cap int
	top int
}

func Init(cap int) *Stack {
	arr := make([]int, cap)
	return &Stack{arr, cap, -1}
}

func (s *Stack) Push(x int) {
	if s.top == s.cap-1 {
		fmt.Println("Stack is Full")
		return
	}
	s.top++
	s.arr[s.top] = x
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		fmt.Println("Underflow")
		return -1
	}
	res := s.arr[s.top]
	s.top--
	return res
}

func (s *Stack) Peek() int {
	if s.top == -1 {
		fmt.Println("Underflow")
		return -1
	}
	return s.arr[s.top]
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
