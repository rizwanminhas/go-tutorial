package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 1-
	fmt.Println("1-first")
	defer fmt.Println("1-second") // executes before it returns the function in the scope i.e. main function in this case.
	fmt.Println("1-third")

	// 2- deferred statements are executed in a FILO order - like a stack
	defer fmt.Println("2-first")
	defer fmt.Println("2-second")
	defer fmt.Println("2-third")

	// 3- We can use defer to close resources, the benefit is that you can put the statement right after you open the resource but it will be executed just before returning from the function.
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)

	// 4- this will print first because it takes the value of the variable at the time when defer is called not at the time when it is executed.
	a := "4-first"
	defer fmt.Println(a)
	a = "4-second"
}
