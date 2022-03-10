package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	// 2nd way
	var responseString strings.Builder
	byteCount, _ := responseString.Write(dataBytes)
	fmt.Println("The Byte Count is", byteCount)
	fmt.Println(responseString.String())

}
