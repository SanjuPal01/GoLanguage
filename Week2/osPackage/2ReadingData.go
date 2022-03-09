package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Reading data from the file")
	fileBytes, err := os.ReadFile("./files.log")
	// fileBytes, err := ioutil.ReadFile("./files.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Content: \n%v\n", string(fileBytes))
}
