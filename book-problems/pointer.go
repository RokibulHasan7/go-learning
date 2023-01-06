package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}
func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	x := 1.5
	square(&x)
	fmt.Printf("%.2f", x)

	// new
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
	// "new" takes a type as an argument, allocate enough memory to fit a value of that type and
	// returns a pointer to it
}
