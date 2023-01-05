package main

import (
	"fmt"
	"math"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	x := make([]int, 0)
	fmt.Printf("Enter the size of Slice: ")
	var sz int
	fmt.Scanf("%d", &sz)
	for i := 0; i < sz; i++ {
		var num int
		fmt.Scanf("%d", &num)
		x = append(x, num)
	}
	smallestNum := math.MaxInt64
	for i := 0; i < sz; i++ {
		smallestNum = min(smallestNum, x[i])
	}
	fmt.Printf("Smallest Number: %d", smallestNum)
	fmt.Println()

	// Slicing from array
	arr := [6]string{"a", "b", "c", "d", "e", "f"}
	sliceFromArray := arr[2:5]
	fmt.Printf("Slice from Array: ")
	fmt.Println(sliceFromArray)

	// Length of Slice
	sl := make([]int, 3, 9)
	fmt.Println("Length of Slice: ", len(sl))
}
