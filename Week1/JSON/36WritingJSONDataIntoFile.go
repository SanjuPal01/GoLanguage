package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./jsonData.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	JSONDataFromWeb := `[
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
	]`
	valid := json.Valid([]byte(JSONDataFromWeb))
	if !valid {
		panic("JSON DATA is not Valid")
	}
	len, err := file.Write([]byte(JSONDataFromWeb))
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is", len)
}
