package main

import "fmt"

func main() {
	// 1- with one variable, i is only available within for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2- with 2 variables
	for i, j := 0, 0; i < 5 && j < 20; i, j = i+1, j+5 {
		fmt.Println(i, j)
	}

	// 3- another weird way to loop, break and continue
	a := 0
	for a < 5 {
		if a == 2 {
			break
		}
		fmt.Println(a)
		a++
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// 4- Gotos are back...? :(

GotoIsBackDude:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i%2 == 0 && j > 3 {
				break GotoIsBackDude
			}
			fmt.Println(i, j)
		}
	}

	// 5- iterate over collections
	b := []int{1, 2, 3}
	for k, v := range b {
		fmt.Println(k, v)
	}

	c := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range c {
		fmt.Println(k, v)
	}

	for k, v := range "I am a good 'ol string" {
		fmt.Println(k, string(v))
	}

}
