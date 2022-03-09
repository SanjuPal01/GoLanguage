package main

import "fmt"

// method is something which we can bound with something
// like we are binding functions with structs so any struct object can access it using . operator

type greeter struct {
	greeting string
	name     string
}

func main() {
	g1 := greeter{
		name:     "Sanju",
		greeting: "Hello",
	}
	g1.greet1()
	fmt.Println("inside main", g1.name) // it didn't changed because it is passed by value. If the data is large in size don't do this instead pass ref

	g2 := greeter{
		greeting: "Hii",
		name:     "Karan",
	}

	g2.greet2()
	fmt.Println("inside main", g2.name)
}

func (g greeter) greet1() {
	fmt.Println(g.greeting, g.name)
	g.name = "Sameer"
	fmt.Println("Inside Function", g.name)
}

func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Sameer"
	fmt.Println("Inside Function", g.name)

}
