package main

import (
	"fmt"
	"reflect"
)

// 1- one way to define a struct (it is like a POJO or case class in scala)
type user struct {
	name  string
	age   int
	email string
}

func main() {
	user1 := user{
		name:  "rizwan",
		age:   100,
		email: "test@go.com",
	}
	fmt.Println(user1)

	// 2- another way to create a struct aka anonymous struct
	user2 := struct {
		name  string
		age   int
		email string
	}{name: "rizwan", age: 100, email: "test@go.com"}
	fmt.Println(user2.name)

	// 3- struct assignment is like arrays i.e. go creates a new struct instead of pointing to same location
	user3 := user1
	user1.name = "minhas" // this will NOT change the name of user3, if you wanted it to also change user3 then use pointers like user3 := &user1. & means address of
	fmt.Println(user1, user3)

	// polymorphism | go uses composition instead of classical inheritance.
	// 4 - you can embed one struct in another.
	type animal struct {
		name   string
		origin string
	}
	type bird struct {
		animal   // bird has animal like behavior...
		canFly   bool
		speedKPH float32
	}

	// 5- you can call on the fields of embedded struct as if it was part of this struct
	bird1 := bird{canFly: true, speedKPH: 50}
	bird1.name = "ostrich" // name is a animal struct property
	fmt.Println(bird1)

	// 6- you can also initialize the embedded struct in the literal syntax
	bird2 := bird{animal: animal{name: "qwe", origin: "asd"}, canFly: true, speedKPH: 40}
	fmt.Println(bird2)

	// 7- Tags - like annotations in JVM based langs...
	type myStruct struct {
		field1 string `required foobar: "100"`
		field2 int
	}

	t := reflect.TypeOf(myStruct{})
	myField1, _ := t.FieldByName("field1")
	fmt.Println(myField1.Tag) // now your code can do anything based on the value of the tag.
}
