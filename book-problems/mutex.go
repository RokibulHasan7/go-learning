package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex

func summ(sl []int, total *int) {
	defer wg.Done()
	m.Lock()
	for _, num := range sl {
		*total += num
	}
	m.Unlock()
	fmt.Printf("Sum: %d", *total)
	//return sumOfSlice
}
func main() {
	//input := bufio.NewReader(os.Stdin)
	//output := bufio.NewWriter(os.Stdout)
	//defer output.Flush()

	n := 20
	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	total := 0
	for i := 0; i < n/5; i++ {
		wg.Add(1)
		go summ(arr[i*5:5*(i+1)], &total)
	}
	wg.Wait()
	fmt.Printf("Sum: %d", total)
}
