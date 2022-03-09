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
        {
                "StudentID": 1,
                "StudentName": "Sameer",
                "tags": [
                        "Physics",
                        "Chemistry",
                        "Biology"
                ]
        }
		`)
	// fmt.Println(string(dataFromWeb))
	isValid := json.Valid(dataFromWeb)
	if isValid {
		// 1st way
		var myData1 Student
		err := json.Unmarshal(dataFromWeb, &myData1)
		if err != nil {
			panic("Can't able to Unmarshal")
		}
		fmt.Println(myData1)

		// 2nd way
		var myData map[string]interface{}
		err = json.Unmarshal(dataFromWeb, &myData)
		if err != nil {
			panic("Can't able to Unmarshal")
		}
		for k, v := range myData {
			fmt.Printf("The Key is %v and the Value is %v and the type is %T\n", k, v, v)
		}

	} else {
		fmt.Println("Not Valid JSON")
	}

}
