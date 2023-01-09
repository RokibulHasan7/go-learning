package main

import (
	"bufio"
	. "fmt"
	"os"
	"sync"
)

func sum(sl []int) int {
	defer wg.Done()
	sumOfSlice := 0
	for _, num := range sl {
		sumOfSlice += num
	}
	return sumOfSlice
}
func main() {
	//input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	wg := new(sync.WaitGroup)
	n := 20
	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	total := 0
	for i := 0; i < n/5; i++ {
		wg.Add(1)
		go sum(arr[i*5 : 5*(i+1)])
	}
	wg.Wait()

	Fprintf(output, "Sum: %d", total)
}
