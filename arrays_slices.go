package main

import "fmt"

/*
1- Arrays has a fixed size that needs to be declared at the compile time.
2- Array assignments or passing them as params causes new array creation. You can use pointers to avoid this.
*/
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

	// 4 - array assignment copies all the contents
	arr4 := [...]int{1, 2, 3}
	arr5 := arr4 // If I did this &arr4 then arr5 would have pointed to the same memory location as arr4
	arr5[1] = 100
	fmt.Println(arr4, arr5)

	// 5 - Slice, can be created by using empty square brackets then type and then initial value. You can also use slice[:], slice[start:], slice[:end] or slice[start:end]
	slice1 := []int{1, 2, 4}
	fmt.Printf("%v, %T, %v\n", slice1, slice1, cap(slice1))

	slice2 := make([]int, 5, 100) // another way of creating the slice, 3=length, 100=capacity
	fmt.Println(slice2)

}
