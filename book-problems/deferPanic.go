package main

import "fmt"

func main() {
	// defer schedules a function call to be run after the function complete
	// f, _ := os.Open(filename)
	// defer f.close()

	// Panic immediately stops the execution of the function
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

	// Recover stop the panic and returns the value that was passed to call to panic.
}