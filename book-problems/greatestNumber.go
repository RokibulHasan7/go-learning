package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func maxNumber(sl ...int) int {
	var mx int
	for _, num := range sl {
		mx = max(mx, num)
	}
	return mx
}
func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println("Greatest number in the list: ", maxNumber(sl...))
}
