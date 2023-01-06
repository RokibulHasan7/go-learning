package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	fmt.Printf("Enter a Number: ")
	var num int
	fmt.Scanf("%d", &num)
	ans := fib(num)
	fmt.Println(ans)
}
