package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./source.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	len, err := io.WriteString(file, "Woo-Hoo\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length is", len)
}
