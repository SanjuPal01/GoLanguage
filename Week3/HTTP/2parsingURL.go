package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://lco.dev:3000/learn?course=golang&paymentid=abcd123"

func main() {
	fmt.Println("Parsing URL")
	result, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Type of URL is %T\n", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The Type of Query Param is %T\n", qParams)
	fmt.Println(qParams["course"])

	for key, value := range qParams {
		fmt.Println(key, ":", value)
	}
}
