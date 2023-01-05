package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	var celsius float64
	celsius = float64((fahrenheit - float64(32)) * float64(float64(5)/float64(9)))
	fmt.Printf("Temperature in Celsius: %.2f", celsius)
}
