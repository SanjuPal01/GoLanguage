package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter Binary: ")
	var binary string
	fmt.Scanln(&binary)
	val, err := binaryToDecimal(binary)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Binary: ", binary)
	fmt.Println("Decimal: ", val)
}
func binaryToDecimal(str string) (int, error) {
	j := len(str) - 1
	decimalVal := 0
	pow := 1
	for j >= 0 {
		if str[j] != '1' && str[j] != '0' {
			return -1, errors.New("Invalid Binary")
		}
		if str[j] == '1' {
			decimalVal += pow
		}
		pow = pow * 2
		j--
	}
	return decimalVal, nil
}
