package main

import "fmt"

func main() {
	// 1st way
	mp := map[int]string{
		1: "Sanju",
		2: "Sameer",
	}
	// 2nd way
	mp3 := make(map[string]string)

	mp3["sanju"] = "1"
	fmt.Println((mp3))
	fmt.Println(mp)
	mp[3] = "Karan"
	fmt.Println(mp)
	mp2 := mp //pass by ref
	mp2[4] = "abc"
	fmt.Println(mp)
	fmt.Println(mp2)
	fmt.Println(mp[2])
	val, ok := mp[1]
	fmt.Println(val, ok)
	val2, ok2 := mp[5]
	fmt.Println(val2, ok2)
	fmt.Println(mp)
	delete(mp, 4)
	fmt.Println(mp)
}
