package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("HTTP - POST request")

	const myURL = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"name": "Sanju",
			"age": 21
		}
	`)
	resp, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}
