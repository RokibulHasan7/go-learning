package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Taking multiple line of input
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		line := scn.Text()
		fmt.Println(line)
		if len(line) == 0 {
			break
		}
	}
}
