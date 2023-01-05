package main

import "fmt"

func sum(sl []int) int {
	sumOfSlice := 0
	for i := 0; i < len(sl); i++ {
		sumOfSlice += sl[i]
	}
	return sumOfSlice
}
func main() {
	sl := []int{1, 2, 3, 4, 5}
	ans := sum(sl)
	fmt.Printf("Sum of Slice numbers: %d.", ans)
}
