package main

import "fmt"

type Student struct {
	id       int
	name     string
	subjects []string
}

func main() {
	// 1st way
	obj1 := Student{
		id:   03,
		name: "Sanju",
		subjects: []string{
			"ISM",
			"SC",
			"IP",
		},
	}
	fmt.Println(obj1)

	// 2nd way - Not Preffered
	obj2 := Student{
		01,
		"Sameer",
		[]string{
			"abc",
			"def",
		},
	}
	fmt.Println(obj2)

	// 3rd way - Anonymous
	obj3 := struct {
		id   int
		name string
	}{id: 06, name: "demo"}
	fmt.Println(obj3)

	// struct follow copy approach - like arrays not slices
	obj4 := obj1
	obj4.id = 45
	fmt.Println(obj4)
	fmt.Println(obj1)

	// for passing by ref - use '&'
	obj5 := &obj1
	obj5.id = 45
	fmt.Println(obj5)
	fmt.Println(obj1)

}
