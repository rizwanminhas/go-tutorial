package main

import (
	"fmt"
)

func main() {
	myBool := true
	// 1-
	if myBool {
		fmt.Println("1.")
	}

	// 2-
	myMap := map[string]int{
		"one": 1,
	}

	if pop, ok := myMap["two"]; ok {
		fmt.Println(pop, "exists")
	}

	// 3- switch statements
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// 4- select with multiple case values
	switch 2 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("other")
	}

	// 5- select with case ranges and fallthrough
	a := 10
	switch {
	case a <= 10:
		fmt.Println("<= 10")
		fallthrough // you can use fallthrough to go to the rest of the cases, golang has implicit "break"
	case a <= 20:
		fmt.Println("<= 20")
	default:
		fmt.Println("defaullt")
	}

	// 6- select to detect type
	var i interface{} = 3
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}

}
