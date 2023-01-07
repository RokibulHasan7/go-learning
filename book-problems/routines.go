package main

import (
	"fmt"
	"time"
)

func say(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf(str)
		fmt.Println()
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	go say("Hey") // Go routine
	go say("Rakib")
	// If we make both "say" Go routines then the program finish before
	// both Go routine finish their task
	time.Sleep(time.Millisecond * 500) // This will make the program wait
}
