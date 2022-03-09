package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rect) Area() float64 {
	return r.height * r.width
}

func main() {
	c := circle{4.5}
	r := rect{5, 7}
	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
