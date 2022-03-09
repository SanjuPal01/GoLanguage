package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Writing data into the File")
	writeFile("./myFile.txt")
}

func writeFile(filename string) {
	// Creating the File
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	content := "Hey Sanju here"
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", len)
}
