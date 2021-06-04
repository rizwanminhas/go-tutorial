package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	// 1-
	ch1 := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch1
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch1 <- 100
		wg.Done()
	}()
	wg.Wait()

	// 2- we can have functions which can only read from channel or write to channel, we can put that constraint using function signature:

	// func myFunc(ch <-chan int) // this function can only read from an int channel
	// func myFunc(ch chan<- string) // this function can only write to a string channel

	// 3- To add buffer to a channel use make(chan int, 50) // this channel can store up to 50 ints without requiring anyone to read from it.

	// 4- you close a channel using close(ch) // but you can't know if a channel is closed or not, using a closed channel will cause panic.

	// 5- you can use for range to loop over channel for i := range ch { ... }
}
