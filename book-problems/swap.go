package main

import "fmt"

func swap(number1 *int, number2 *int) {
	tmp := *number1
	*number1 = *number2
	*number2 = tmp
}
func main() {
	fmt.Printf("Enter 2 number: ")
	var number1, number2 int
	fmt.Scanf("%d %d", &number1, &number2)
	fmt.Printf("Number1: %d, Number2: %d \n", number1, number2)
	swap(&number1, &number2)
	fmt.Printf("Number1: %d, Number2: %d", number1, number2)
}
