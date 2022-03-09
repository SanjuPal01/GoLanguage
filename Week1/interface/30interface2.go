package main

import "fmt"

type Detail1 interface {
	getID() int
	getName() string
}

type Detail2 interface {
	getLoc() string
}

type Detail interface {
	Detail1
	Detail2
}

type student struct {
	stuId   int
	stuName string
	stuLoc  string
}

type faculty struct {
	facId   int
	facName string
	facLoc  string
}

func (s student) getID() int {
	return s.stuId
}

func (s student) getName() string {
	return s.stuName
}

func (s student) getLoc() string {
	return s.stuLoc
}

func (f faculty) getID() int {
	return f.facId
}

func (f faculty) getName() string {
	return f.facName
}

func (f faculty) getLoc() string {
	return f.facLoc
}

func main() {
	s := student{stuId: 03, stuName: "Sanju", stuLoc: "Chandigarh"}
	f := faculty{facId: 01, facName: "Sameer", facLoc: "Chandigarh"}
	d := []Detail{s, f}
	for _, detail := range d {
		fmt.Printf("%v, %v, %v\n", detail.getID(), detail.getName(), detail.getLoc())
	}

	/*
		d2 := []Detail{s, f}
		for _, detail := range d2 {
			fmt.Printf("%v\n", detail.getLoc())
		}
	*/
}
