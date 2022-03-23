package main

import "fmt"

func comparison(arr1, arr2 []int) {
	size1 := len(arr1)
	size2 := len(arr2)
	// Storing second array elemets - so i can check with arr1 elements is it present in arr2 or not
	myMap := map[int]int{}
	for i := 0; i < size2; i++ {
		myMap[arr2[i]]++
	}
	for i := 0; i < size1; i++ {
		_, ok := myMap[arr1[i]]
		if !ok {
			fmt.Printf("This Number %d is Not present in Second Array\n", arr1[i])
		}
	}
}

func main() {
	var size int
	fmt.Print("Enter 1st Array Size: ")
	fmt.Scanln(&size)
	arr1 := make([]int, size)
	fmt.Println("Enter Array Elements ")
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %d Element: ", i+1)
		fmt.Scanln(&arr1[i])
	}
	fmt.Print("Enter 2nd Array Size: ")
	fmt.Scanln(&size)
	arr2 := make([]int, size)
	fmt.Println("Enter Array Elements")
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %d Element: ", i+1)
		fmt.Scanln(&arr2[i])
	}
	comparison(arr1, arr2)
}
