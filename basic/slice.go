package main

import "fmt"

func main() {
	// Declaration way of Slice
	slice1 := make([]int, 5)
	fmt.Println(slice1) // [0 0 0 0 0]

	var slice2 []int
	fmt.Println(slice2) // []

	slice3 := []int{1, 2, 3}
	fmt.Println(slice3) // [1, 2, 3]

	// Create a slice from an array
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr1[1:3]
	fmt.Println(slice4) // [2, 3]

	// Append
	slice5 := []int{1, 2, 3}
	slice5 = append(slice5, 4, 5)
	fmt.Println(slice5) // [1, 2, 3, 4, 5]

	// Append multiple slice
	mySlice1 := []int{1, 2, 3}
	mySlice2 := []int{4, 5}
	mySlice3 := append(mySlice1, mySlice2...)
	fmt.Println(mySlice3) // [1, 2, 3, 4, 5]
	mySlice3 = append(mySlice3, mySlice3...)
	fmt.Println(mySlice3) // [1, 2, 3, 4, 5, 1, 2, 3, 4, 5]

	// Copy
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	secondSlice := originalSlice[:len(originalSlice)-10]
	fmt.Println(secondSlice) // [1,2,3,4,5]
	copySlice := make([]int, len(secondSlice))
	copy(copySlice, secondSlice) // copy(dest, src)
	fmt.Println(copySlice)       // [1,2,3,4,5]
}
