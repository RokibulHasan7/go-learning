package main

import "fmt"

func main() {
	for i := 3; i < 101; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	i := 3
	for i < 101 {
		fmt.Printf("%d ", i)
		i += 3
	}
}