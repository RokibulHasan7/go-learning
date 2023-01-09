package math

import sm "math"

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	mn := sm.MaxFloat64
	for _, x := range xs {
		if mn > x {
			mn = x
		}
	}
	return mn
}

func Max(xs []float64) float64 {
	mx := sm.SmallestNonzeroFloat64
	for _, x := range xs {
		if mx < x {
			mx = x
		}
	}
	return mx
}
