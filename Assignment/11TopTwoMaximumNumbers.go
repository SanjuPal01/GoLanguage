package main

import (
	"fmt"
	"math"
)

func TopTwoUnique(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
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
		return []int{first, second}
	} else {
		return []int{-1}
	}
}
func TopTwoRepetition(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
	}
	first := math.MinInt
	second := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > first {
			second = first
			first = arr[i]
		} else if arr[i] > second {
			second = arr[i]
		}
	}
	if second != math.MinInt {
		return []int{first, second}
	} else {
		return []int{-1}
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
	res := TopTwoUnique(arr)
	if res[0] == -1 {
		fmt.Println("No Top Two Elements")
	} else {
		fmt.Println("Top Two Unique Highest Element is:", res[0], res[1])
	}

	// res = TopTwoRepetition(arr)
	// if res[0] == -1 {
	// 	fmt.Println("No Top Two Elements")
	// } else {
	// 	fmt.Println("Top Two Repeated Highest Element is:", res[0], res[1])
	// }
}
