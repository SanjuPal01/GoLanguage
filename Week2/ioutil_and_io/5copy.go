package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("./source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.Create("./destination.txt")
	len, err := io.Copy(dest, src)
	defer dest.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length is", len)
}
