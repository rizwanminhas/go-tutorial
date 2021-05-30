package main

import "fmt"

/*
1- an array is a valid key type but a slice is not. map[[3]int]string{} is ok but map[[]int]string{} is not ok.
2- order of keys in a map is not guaranteed.
3- assigning map to another variable just points to the original so mutating one will impact the other (just like the slice)
*/
func main() {

	// 1- first way to create a map
	myMap1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap1)

	// 2- second way to create a map
	myMap2 := make(map[string]int)
	myMap2["one"] = 1
	fmt.Println(myMap2)

	// 3- to delete a key use this
	delete(myMap1, "two") // deleting a non existing key just silently returns
	fmt.Println(myMap1)

	// 4- non existing key just returns the default for the value type
	fmt.Println(myMap1["foo"])

	// 5- use this to check if a value exists or not
	pop, exists := myMap1["foo"]
	fmt.Println(pop, exists)

	// 6- use len to check elems count
	fmt.Println(len(myMap1))

}
