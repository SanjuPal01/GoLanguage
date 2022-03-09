package main

import "fmt"

func main() {
	// 1st way
	arr1 := []int{1, 2, 3} //it's slice. You can append element in this. You can give ref to another variable by simple assignment
	fmt.Println(arr1)
	a := arr1 // by ref
	a[2] = 5
	fmt.Println(arr1, a)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	arr1 = append(arr1, 4) //can append data in slice
	fmt.Println(arr1)

	// 2nd way
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := arr2[:]
	c := arr2[3:]
	d := arr2[:6]
	e := arr2[3:6]
	fmt.Println(arr2, b, c, d, e)
	arr2[5] = 12
	fmt.Println(arr2, b, c, d, e)

	// 3rd way
	arr3 := make([]int, 3) //type and then followed by size
	fmt.Println(arr3)
	arr4 := make([]int, 3, 100) //type , size and then cap
	fmt.Println((arr4))

	// 4th way - Blank slice
	arr5 := []int{}
	fmt.Println(arr5)
	fmt.Println(len(arr5))
	fmt.Println(cap(arr5))
	arr5 = append(arr5, 1)
	fmt.Println(arr5)
	fmt.Println("Length: ", len(arr5))
	fmt.Println("Capacity: ", cap(arr5))
	arr5 = append(arr5, 2, 3, 4, 5)
	fmt.Println(arr5)
	fmt.Println("Length: ", len(arr5))
	fmt.Println("Capacity: ", cap(arr5))

	// If you want to concat two slice then we need to use spread operator - "..."
	arr5 = append(arr5, []int{6, 7, 8, 9}...)
	fmt.Println(arr5)

	// If you want to concat two slices and make one new slice then we need to do this
	t1 := []int{1, 2}
	t2 := []int{3, 4, 5}
	arr6 := []int{}
	arr6 = append(arr6, t1...)
	arr6 = append(arr6, t2...)
	fmt.Println(arr6)

	// Another way to do the above
	arr7 := append(t1, t2...)
	fmt.Println(arr7)

	// Getting first character removed, then last character and then any character in between
	fmt.Println("Original Array: ", arr6)
	firstRemoved := arr6[1:]
	fmt.Println("After First Removed: ", firstRemoved)
	lastRemoved := arr6[:len(arr6)-1]
	fmt.Println("After Last Removed: ", lastRemoved)
	middleRemoved := append(arr6[:2], arr6[3:]...)
	fmt.Println("After Middle Removed: ", middleRemoved)

}
