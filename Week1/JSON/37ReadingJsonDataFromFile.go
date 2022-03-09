package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	ID       int      `json:"StudentID"`
	Name     string   `json:"StudentName"`
	Courses  []string `json:"tags"`
	Password string   `json:"-"`
}

func main() {
	dataByte, err := ioutil.ReadFile("./jsonData.json")
	if err != nil {
		panic(err)
	}
	// JSONData := string(dataByte)
	// fmt.Println(JSONData)

	//1st way
	var content []Student
	err = json.Unmarshal(dataByte, &content)
	if err != nil {
		panic(err)
	}
	for i := range content {
		fmt.Println(content[i])
	}
	fmt.Println()

	// 2nd way
	var content2 []map[string]interface{}
	err = json.Unmarshal(dataByte, &content2)
	if err != nil {
		panic(err)
	}
	// fmt.Println(content2)
	for i := range content2 {
		for k, v := range content2[i] {
			fmt.Printf("The key is %v and the Value is %v and the Type is %T\n", k, v, v)
		}
		fmt.Println()
	}
}
