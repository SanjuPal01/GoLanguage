package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hey Everybody, Sanju Here, How's everyone doing?"
	//Splitting
	splittedString := strings.Split(str, ",")
	fmt.Printf("%v, %T\n", splittedString, splittedString)
	fmt.Println()

	for _, v := range splittedString {
		fmt.Println(v)
	}

	// Joining
	joinedString := strings.Join(splittedString, "-->")
	fmt.Println(joinedString)
}
