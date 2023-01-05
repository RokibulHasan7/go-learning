package main

import (
	"fmt"
)

func half(num int) (int, bool) {
	halfOfNum := num / 2
	//fmt.Println(halfOfNum)
	if halfOfNum == num {
		return halfOfNum, true
	} else {
		return halfOfNum, false
	}
}
func main() {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num)
	ans, even := half(num)
	if even == false {
		fmt.Printf("Odd, Half of Number: %d", ans)
	} else {
		fmt.Printf("Even, Half of Number: %d", ans)
	}
}
