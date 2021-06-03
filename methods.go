package main

import "fmt"

/*
Methods are functions that run in a given context
*/

type greeter struct {
	name     string
	greeting string
}

// 1- greeter is the context and greet is the function name. It will get a copy of the greeter object so changes made here won't affect the caller's greeter.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {

	// 1-
	g := greeter{
		name:     "rizwan",
		greeting: "hello!",
	}

	g.greet() // NOTE: greet method is available on greeter struct.
}
