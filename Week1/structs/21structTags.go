package main

import (
	"fmt"
	"reflect"
)

// tags will help you to give meta info - which you can access using reflection package
type Student struct {
	id   int `required max:10`
	name string
}

func main() {
	t := reflect.TypeOf(Student{})
	field, _ := t.FieldByName("id")
	fmt.Println(field.Tag)
}
