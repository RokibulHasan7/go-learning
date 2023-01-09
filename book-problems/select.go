package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from c1"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			c2 <- "from c2"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1) // print "from c1" every 100 milisec
			case msg2 := <-c2:
				fmt.Println(msg2) // print "from c2" every 200 milisec
			case <-time.After(time.Millisecond * 500): // time.After creates a channel and after given time it will
				// the current time on it.
				fmt.Println("timeout")
			default:
				//fmt.Println("No channel ready.") // If no channels are ready
			}
		}
	}()
	time.Sleep(time.Millisecond * 10) // To delay the program
}
