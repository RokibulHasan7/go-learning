package main

import (
	"bufio"
	. "fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func sum(ch chan int, sl []int) {
	defer wg.Done()
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
	ch := make(chan int, 4)
	/*Fprintf(output, "Enter %d numbers: ", n)
	for i := 0; i < n; i++ {
		Fscanf(input, "%d", &arr[i])
	}
	/*wg.Add(1)
	go sum(ch, arr[0:5])

	wg.Add(1)
	go sum(ch, arr[5:10])

	wg.Add(1)
	go sum(ch, arr[10:15])

	wg.Add(1)
	go sum(ch, arr[15:20])*/

	for i := 0; i < n/5; i++ {
		wg.Add(1)
		go sum(ch, arr[i*5:5*(i+1)])
	}
	wg.Wait()

	close(ch)
	total := 0
	for item := range ch {
		total += item
	}
	Fprintf(output, "Sum: %d", total)
}
