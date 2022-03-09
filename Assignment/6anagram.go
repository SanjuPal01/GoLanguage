package main

import "fmt"

func main() {
	fmt.Print("Enter string1: ")
	var str1, str2 string
	fmt.Scanln(&str1)
	fmt.Print("Enter string2: ")
	fmt.Scanln(&str2)
	if isAnagram(str1, str2) {
		fmt.Printf("%s and %s are anagram\n", str1, str2)
	} else {
		fmt.Printf("%s and %s are not anagram\n", str1, str2)
	}
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	count := make([]int, 256)
	for i := 0; i < len(str1); i++ {
		// fmt.Println(int(str1[i]))
		count[int(str1[i])]++
		count[int(str2[i])]--
	}
	for i := 0; i < 256; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
