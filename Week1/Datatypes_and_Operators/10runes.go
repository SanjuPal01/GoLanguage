// It is similar to strings(but can contain only single character) and it uses Unicode32-UTF32
// we use '' while defining it
// see the comparison between string and rune for clear understanding
package main

import "fmt"

func main() {
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var s rune = 'a'
	fmt.Printf("%v, %T\n", s, s)
}
