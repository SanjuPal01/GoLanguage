package main

import "fmt"

func main() {
	fmt.Print("Enter Start and End Value: ")
	var start, end int
	fmt.Scanln(&start, &end)
	total, err := sumOf(start, end)
	if err != nil {
		fmt.Println("There is an error:", err)
	} else {
		fmt.Println(total)
	}
}
func sumOf(start, end int) (int, error) {
	if start > end {
		return 0, fmt.Errorf("Start is greater than End (%d > %d)", start, end)
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
