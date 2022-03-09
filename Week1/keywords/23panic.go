package main

import "fmt"

// it is similar to error. we use it where we need to stop our program. After that your program will be terminated
// First normal statements will execute then defer statements, and then panic statements.
func main() {
	fmt.Println("Start")
	defer fmt.Println("Checking")
	panic("Middle")
	fmt.Println("end") //Never reach
}
