package main

import (
	"fmt"
	"unicode"
)

func main() {
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	fmt.Printf("%v %T\n", mixed, mixed)
	fmt.Printf("%q %T\n", mixed[len(mixed)-1], mixed[0])
	for _, c := range mixed {
		fmt.Printf("For Value: %q and Type: %T:\n", c, c)
		if unicode.IsControl(c) {
			fmt.Println("\tis Control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis Digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}
}
