package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var testcase int
	Fscan(input, &testcase)
	for t := 0; t < testcase; t++ {
		var str string
		Fscan(input, &str)
		//Fprintln(output, str)
		sz := len(str)
		pos := 0
		for i := 1; i < sz - 1; i++ {
			if str[i] == 'a' {
				pos = i
				break
			}
		}
		//Fprint(output, pos, "\n")
		if pos != 0 {
			Fprint(output, str[0:pos], " ", "a ", str[pos+1:], "\n")
		} else {
			Fprint(output, string(str[0]), " ", str[1:sz-1], " ", string(str[sz-1]), "\n")
		}
	}
}
