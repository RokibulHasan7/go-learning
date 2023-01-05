package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meter := float64(feet * 0.3048)
	fmt.Printf("In meter: %.2f", meter)
}
