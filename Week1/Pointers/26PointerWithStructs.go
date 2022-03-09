package main

import "fmt"

type Student struct {
	id int
}

func main() {
	stu1 := Student{id: 23}
	fmt.Println(stu1)

	var stuPointer *Student
	fmt.Println(stuPointer) //points to nil. Hence you can't dereference it using *
	stuPointer = &stu1

	// 1st way
	fmt.Println((*stuPointer).id) //first use () then . - because . has more precedence than * otherwise it would be treated as *(stuPointer.id) and it's just a pointer hence it won't have id field
	(*stuPointer).id = 24
	fmt.Println((*stuPointer).id)

	//2nd way - Go lang with do the above things by ourself. Easy to read but hard to understand for programmers who coming from c/c++
	fmt.Println(stuPointer.id)
	stuPointer.id = 26
	fmt.Println(stuPointer.id)

	// new keyword
	stuPointer2 := new(Student) // Create blank struct - Can't initialize the fields
	fmt.Println(*stuPointer2)
	(*stuPointer2).id = 02
	fmt.Println(*stuPointer2)

}
