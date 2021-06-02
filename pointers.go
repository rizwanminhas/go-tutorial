package main

import "fmt"

func main() {

	// 1-
	a1 := 10
	var b1 *int = &a1    // *int means "pointer to an integer", &a1 means "address of a1"
	fmt.Println(a1, b1)  // b1 will contain the memory address of a1
	fmt.Println(a1, *b1) // *b1 means "dereference", *b1 will now print the value at the memory location stored in b1

	// 2-
	type myStruct struct {
		foo int
	}
	var a2 *myStruct // default value of reference is <nil>
	fmt.Println(a2)

	// 3-
	var a3 *myStruct = new(myStruct)
	a3.foo = 10 // same as (*a3).foo = 10
	fmt.Println(a3, *a3)
}
