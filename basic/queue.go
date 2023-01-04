package main

import "fmt"

func push(queue []int, num int) []int {
	queue = append(queue, num)
	fmt.Println("After Push: ", queue)
	return queue
}

func pop(queue []int) []int {
	queue = queue[1:]
	fmt.Println("After Pop: ", queue)
	return queue
}

func main() {
	var queue []int

	queue = push(queue, 1)
	queue = push(queue, 2)
	queue = push(queue, 3)
	queue = pop(queue)
	queue = push(queue, 4)
}
