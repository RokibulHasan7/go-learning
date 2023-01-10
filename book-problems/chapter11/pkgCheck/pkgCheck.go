package main

import (
	"example.com/chapter11/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	mn := math.Min(xs)
	mx := math.Max(xs)
	fmt.Printf("Min number: %.0f, Max number: %.0f", mn, mx)
}
