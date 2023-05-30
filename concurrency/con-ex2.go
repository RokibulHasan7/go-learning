package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go hw(&wg)
	go gb(&wg)
	wg.Wait()
}

func hw(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World!")
}

func gb(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye!")
}
