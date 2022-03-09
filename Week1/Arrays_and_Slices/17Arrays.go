package main

import "fmt"

func main() {

	// Arrays
	// in golang array don't share the ref variable by default if you do a=b then b array is copied to a.

	// 1st way
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 2nd way
	arr2 := [...]string{"sanju", "pal"}
	fmt.Println(arr2)

	// 3rd Way
	var arr3 [3]int
	fmt.Println(arr3)
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3
	//arr3 = append(arr3, 5)		You can't append in the array(for this you need to use slice)
	fmt.Println(arr3)
	fmt.Println(arr3[1])
	fmt.Println("Length: ", len(arr3))
	fmt.Println("Capacity: ", cap(arr3))

	b := arr3
	b[2] = 5
	// You can clearly see here the array is not sending its reference value it's just copying it
	fmt.Println("See the change here", arr3, b)

	// pass arr by ref (Pointers concept)
	c := &arr3
	c[2] = 6
	fmt.Println("See the change here", arr3, c)

	// 2d array
	// 1st way - hard to read
	var identityMatrix1 [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix1)

	// 2nd way
	var identityMatrix2 [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix2)

	// 3rd way
	var identityMatrix3 [3][3]int
	identityMatrix3[0] = [3]int{1, 0, 0}
	identityMatrix3[1] = [3]int{0, 1, 0}
	identityMatrix3[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix3)
}
