package main

import (
	"fmt"
	"math"
)

func SecondHighest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	first := math.MinInt
	second := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > first {
			second = first
			first = arr[i]
		} else if arr[i] > second && arr[i] < first {
			second = arr[i]
		}
	}
	if second != math.MinInt {
		return second
	} else {
		return -1
	}
}

func main() {
	var size int
	fmt.Print("Enter 1st Array Size: ")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("Enter Array Elements ")
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %d Element: ", i+1)
		fmt.Scanln(&arr[i])
	}
	res := SecondHighest(arr)
	if res == -1 {
		fmt.Println("There is No second Highest Element")
	} else {
		fmt.Println("Second Highest Element is:", res)
	}
}
