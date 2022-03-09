package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Reading Data from file")
	ReadFile("./myFile.txt")
}

func ReadFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // Give data in bytes
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data: %v\n", string(dataByte))
}
