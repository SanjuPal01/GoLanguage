package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print("Enter Start and End: ")
	var start, end int
	fmt.Scanln(&start, &end)
	total, err := sumOf(start, end)
	if err != nil {
		fmt.Println("There is an error:", err)
	} else {
		fmt.Println(total)
	}
}

func sumOf(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("Start is greater than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
