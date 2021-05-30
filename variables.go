package main

import (
	"fmt"
	"strconv"
)

/*
	If a variable or constant name starts with upper case then it is public.
*/
func main() {
	// 1 - declare variable then assign value, if no  value is assigned then the type's default will be assigned e.g. 0 for int, false for boolean
	var a int
	a = 10
	fmt.Printf("a: %v, type=%T\n", a, a)

	// 2 - declare and assign in one line
	var b float32 = 20.0
	fmt.Printf("b: %v, type=%T\n", b, b)

	// 3 - declare and assign short syntax - type will be infered.
	c := "I am a string"
	fmt.Printf("%v, type=%T\n", c, c)

	// 4 - define multiple vars in a var block
	var (
		var1 string = "I am a string"
		var2 bool   = true
		var3 byte   = 'R' // to print 'R' use string(var3)
	)
	fmt.Printf("(%v | %T), (%v | %T), (%v | %T)\n", var1, var1, var2, var2, string(var3), var3)

	// 5 - Following will print the ascii value at 100 i.e. the character `d`
	// var s string = string(100)
	// fmt.Print(s)

	// 6 - To actually print the int 100 use the following
	var s string = strconv.Itoa(100)
	fmt.Printf("%v, %T\n", s, s)

	// 7 - To define a constant use the keywoard const
	const d int = 10
	fmt.Println(d)
}
