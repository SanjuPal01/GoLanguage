package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Creating file and then writing content into it...")
	file, err := os.Create("./files.log")
	if err != nil {
		log.Fatal("Can't able to create file")
	}
	defer file.Close()
	len, err := file.WriteString("Hey Sanju Here")
	if err != nil {
		log.Fatal("Can't able to write the String into the file")
	}
	fmt.Printf("Length of content %d\n", len)
}
