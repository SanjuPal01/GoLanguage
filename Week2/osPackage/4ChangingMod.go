package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chmod("./sanju.log", 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mod is changed successfully")
	// For checking what you can do go and reveal the log file in finder and then get the info of file and see sharing and permission option in the last tab
	// 0644 means - R&W, R, R
	// 0666 means - R&W, R&W, R&W
}
