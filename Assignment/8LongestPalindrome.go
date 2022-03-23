package main

import (
	"fmt"
)

func LongestPalin(str string) {
	n := len(str)

	table := make([][]bool, n)
	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
	}

	// substrings of length 1 are already palindromes
	maxLength := 1

	for i := 0; i < n; i++ {
		table[i][i] = true
	}

	// Finding sub-string of length 2.
	start := 0
	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			table[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i + k - 1
			if table[i][j-1] && str[i] == str[j] {
				table[i][j] = true
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
		}
	}
	fmt.Println("Longest Palindrome Substring is:", str[start:start+maxLength])
}

func main() {
	var str string
	fmt.Print("Enter Input: ")
	fmt.Scanln(&str)
	LongestPalin(str)
}
