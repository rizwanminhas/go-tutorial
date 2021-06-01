package main

import (
	"fmt"
)

func main() {
	// 1-
	// fmt.Println("1-first")
	// panic("oops...")
	// fmt.Println("1-first")

	// 2- a very basic web server
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("rizwan minhas"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// 3- panic is exectued after defer
	// fmt.Println("3-first")
	// defer fmt.Println("3-this was defered.")
	// panic("3-oops...")
	// fmt.Println("3-first")

	// 4- Note the main function continues to executes even after calling the panicker.
	fmt.Println("4-start")
	panicker()
	fmt.Println("4-end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
			// panic(err) // re throw the panic if you don't know how to recover from it, this rethrowing of panic will also break the execution in the caller function. i.e. "4-end" will not be printed
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
