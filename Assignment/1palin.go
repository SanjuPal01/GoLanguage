package main

import "fmt"

func main() {
	fmt.Print("Enter String: ")
	var str string
	fmt.Scanln(&str)
	if isPalin(str) {
		fmt.Println(str, "is Palindrome")
	} else {
		fmt.Println(str, "is not a Palindrome")
	}
}

func isPalin(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
