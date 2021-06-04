package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg2 = sync.WaitGroup{}
var mutex = sync.RWMutex{}
var counter = 0

func main() {
	// 1-
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// 2-
	var a1 = "hello"
	go func() {
		fmt.Println(a1)
	}()
	a1 = "value updated" // it will print this because the main thread will change the value before the goroutine can print it.
	time.Sleep(100 * time.Millisecond)

	// 3- wg is used to avoid using the sleep
	var wg = sync.WaitGroup{}
	var a2 = "hello"
	wg.Add(1)
	go func(a2 string) {
		fmt.Println(a2)
		wg.Done()
	}(a2)
	a2 = "value updated"
	wg.Wait()

	// 4- mutex

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg2.Add(2)
		mutex.RLock()
		go sayHello2()
		mutex.Lock()
		go increment()
	}
	wg2.Wait()

	// 5- how many os threads are available?
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// 6- to debug race conditions use
	// go run -race src/<filename>.go

}

// 1-
func sayHello() {
	fmt.Println("hello")
}

func sayHello2() {
	fmt.Printf("hello # %v\n", counter)
	mutex.RUnlock()
	wg2.Done()
}

func increment() {
	counter++
	mutex.Unlock()
	wg2.Done()
}
