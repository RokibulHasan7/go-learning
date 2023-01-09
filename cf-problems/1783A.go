package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var testcase int
	Fscan(input, &testcase)
	for _t := 0; _t < testcase; _t++ {
		var n int
		Fscan(input, &n)
		sl := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(input, &sl[i])
		}

		cnt := 0
		for i := 1; i < n; i++ {
			if sl[i-1] == sl[i] {
				cnt++
			}
		}
		if cnt == n-1 {
			Fprintln(output, "NO")
		} else {
			Fprintln(output, "Yes")

			sort.Slice(sl, func(a, b int) bool {
				return a < b
			})
			for i := 0; i < n/2; i++ {
				Fprintf(output, "%d %d ", sl[i], sl[n-1-i])
			}
			if n%2 == 1 {
				Fprintf(output, "%d", sl[n/2])
			}
			Fprintln(output)
		}
	}

}
