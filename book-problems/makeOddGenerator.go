package main

import "fmt"

func oddGenerator() func() int {
	num := 1
	return func() int {
		odd := num
		num += 2
		return odd
	}
}
func main() {
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
