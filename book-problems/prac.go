package main

import (
	"bufio"
	. "fmt"
	"os"
)

func summ(ch chan int, sl []int) {
	sumOfSlice := 0
	for _, num := range sl {
		sumOfSlice += num
	}
	ch <- sumOfSlice
}
func main() {
	//input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	n := 20
	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ch := make(chan int, n/5)
	total := 0
	for i := 0; i < n/5; i++ {
		go summ(ch, arr[i*5:5*(i+1)])
		total += <-ch
	}

	close(ch)
	Fprintf(output, "Sum: %d", total)
}
