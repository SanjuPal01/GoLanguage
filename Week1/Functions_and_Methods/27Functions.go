package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	greet("Hello")
	sum1 := sum1(2, 3)
	fmt.Println("Sum of 2 and 3 is", sum1)
	sum2 := sum2(5, 7)
	fmt.Println("Sum of 5 and 7 is", sum2)
	sum3 := sum3(5, 4)
	fmt.Println("Sum of 5 and 7 is", *sum3)

	// returning more than one value
	div, err := division(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Division Result", div)

	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of All", sum)

	// Anonymous functions
	// Not a good way using the outer value inside other function without passing it as a parameter - You will see the consequences in go routines
	for i := 0; i < 5; i++ {
		func() {
			fmt.Printf("Hii %d Time\n", i+1)
		}()
	}

	// Good way
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Printf("Hii %d Time\n", i+1)
		}(i)
	}

	// Function as a variable
	// var f func() = func(str string){}	//1st way
	f := func(str string) {
		fmt.Println("hello", str)
	}
	f("Sanju")

	// Passing Function as a parameter - Hard to read
	fun(f)

	// passing function in variable which also return something
	var f2 func(float64, float64) (float64, error)
	f2 = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, errors.New("Divide by Zero")
		}
		return a / b, nil
	}
	res, err := f2(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Division result: ", res)
}

func test() {
	fmt.Println("TESTING")
}

func greet(str string) {
	fmt.Println(str)
}

// 1st way
func sum1(a int, b int) int { //You can also return the local variables - it will pass this variable from stack memory to heap memory automatically
	res := a + b
	return res
}

// 2nd way
func sum2(a, b int) (res int) { // Or you can add the variable name here which you want to return
	res = a + b
	return
}

// 3rd way
func sum3(a, b int) *int { //Returning Pointer
	res := a + b
	return &res
}

func division(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("Divide by Zero")
	}
	return a / b, nil
}

func sumAll(val ...int) int {
	sum := 0
	for _, v := range val {
		sum += v
	}
	return sum
}

func fun(f func(str string)) {
	f("Sameer")
}
