package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./sanju.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("New Line added\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Appended or created new file")
}
