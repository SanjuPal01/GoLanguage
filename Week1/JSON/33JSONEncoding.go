package main

import (
	"encoding/json"
	"fmt"
)

// You can also change the Json key data by adding the tags here
// if you don't want to show password field simply use '-' tag

type student struct {
	Id       int      `json:"StudentID"`
	Name     string   `json:"StudentName"`
	Password string   `json:"-"`
	Courses  []string `json:"tags,omitempty"` // if there is no courses values then it will not be displayed
}

func main() {
	students := []student{
		{Id: 01, Name: "Sameer", Password: "abc123", Courses: []string{"Physics", "Chemistry", "Biology"}},
		{Id: 02, Name: "Karan", Password: "def123", Courses: []string{"Mysql", "ReactJS"}},
		{Id: 03, Name: "Sanju", Password: "pqr123", Courses: []string{"go-lang", "Angular"}},
		{Id: 04, Name: "Ajay", Password: "xyz123"},
	}

	// finalData, err := json.Marshal(students)
	finalData, err := json.MarshalIndent(students, "", "\t") // it will give the data in indented way
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalData)
}
