package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID       int      `json:"StudentID"`
	Name     string   `json:"StudentName"`
	Password string   `json:"-"`
	Courses  []string `json:"tags,omitempty"`
}

func main() {
	dataFromWeb := []byte(`
	[
        {
                "StudentID": 1,
                "StudentName": "Sameer",
                "tags": [
                        "Physics",
                        "Chemistry",
                        "Biology"
                ]
        },
        {
                "StudentID": 2,
                "StudentName": "Karan",
                "tags": [
                        "Mysql",
                        "ReactJS"
                ]
        },
        {
                "StudentID": 3,
                "StudentName": "Sanju",
                "tags": [
                        "go-lang",
                        "Angular"
                ]
        },
        {
                "StudentID": 4,
                "StudentName": "Ajay"
        }
	]
	`)
	// fmt.Println(string(dataFromWeb))
	isValid := json.Valid(dataFromWeb)
	if isValid {
		// 1st way
		var myData1 []Student
		err := json.Unmarshal(dataFromWeb, &myData1)
		if err != nil {
			panic("Can't able to Unmarshal")
		}
		fmt.Println(myData1)
		for _, value := range myData1 {
			fmt.Println(value)
		}
		fmt.Println()
		// 2nd Way
		var myData2 []map[string]interface{}
		err = json.Unmarshal(dataFromWeb, &myData2)
		if err != nil {
			panic("Can't able to Unmarshal")
		}
		for i := range myData2 {
			for k, v := range myData2[i] {
				fmt.Printf("The key is %v and the Value is %v and the Type is %T\n", k, v, v)
			}
			fmt.Println()
		}

	} else {
		fmt.Println("Not Valid JSON")
	}

}
