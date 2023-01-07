package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf(str)
		fmt.Println()
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // make sure the routine done
}
func main() {
	wg.Add(1)     // add before every go routine
	go say("Hey") // Go routine
	wg.Add(1)
	go say("Rakib")
	// If we make both "say" Go routines then the program finish before
	// both Go routine finish their task

	wg.Wait() // this will make the program wait to finish the routine work
	// time.Sleep(time.Millisecond * 500) // This will make the program wait
}
