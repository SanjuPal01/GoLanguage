package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("GET Request")
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("The Response Body Type is %T\n", resp.Body)
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}
