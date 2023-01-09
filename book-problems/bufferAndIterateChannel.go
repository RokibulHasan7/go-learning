package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
	//wg.Done()
}

func main() {
	fooVal := make(chan int, 10) // Add buffer, 10 for 10 go routines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait()     // wait to complete all go routines
	close(fooVal) // Close the channel

	// iterate through the all routine
	for item := range fooVal {
		fmt.Println(item)
	}

}
