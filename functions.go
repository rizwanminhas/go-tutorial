package main

import "fmt"

func main() {
	// 1-
	myFunc1("rizwan", "minhas")

	// 2-
	firstName2 := "rizwan"
	lastName2 := "minhas"
	myFunc2(&firstName2, &lastName2)
	fmt.Println(lastName2 + ", " + firstName2)

	// 3-
	fmt.Printf("Sum1=%v\n", sum1(1, 2, 3, 4, 5))

	// 4-
	sum2 := sum2(1, 2, 3, 4, 5)
	fmt.Printf("Sum2=%v\n", *sum2)

	// 7- lambda and iif
	func() {
		fmt.Println("an imediately invoked anonymous function...")
	}()

	// 8- you can also assign functions to variables
	myFunc3 := func() {
		fmt.Println("function assigned to a variable...")
	}
	myFunc3()
}

// 1- NOTE: since both params had same type we can skip explicitly declaring type of the first param
func myFunc1(firstName, lastName string) {
	fmt.Println("hello1 " + lastName + ", " + firstName)
}

// 2- Since params are reference changing them will effect the caller as well.
func myFunc2(firstName, lastName *string) {
	*firstName = "Riz"
	fmt.Println("hello2 " + *lastName + ", " + *firstName)
}

// 3- variadic aka varargs and a return type
func sum1(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// 4- variadic aka varargs and a return type of a pointer, NOTE: Go will automatically upgrade the local result variable from stack to heap because it recognizes that you are returning a pointer to it
func sum2(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// 5- example of a named return
func sum3(values ...int) (result int) {
	for _, v := range values {
		result += v // the variable result is available to you because you are using the "named return" style.
	}
	return // implicitly returns result
}

// 6- return multiple values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
