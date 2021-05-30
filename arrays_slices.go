package main

import "fmt"

func main() {
	// 1 - Define size, type and initial values
	arr1 := [5]int{} // or to initialize with values use [5]int{1,2,3,4,5}
	fmt.Println(arr1)

	// 2 - use the spread operator to avoid giving the size, size will be fixed to the length of initial values i.e. 3 in this case
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	// 3 - define array variable and the set values later
	var arr3 [3]string
	fmt.Println(arr3)
	arr3[1] = "rizwan"
	fmt.Println(arr3)
	fmt.Println(len(arr3))
}
