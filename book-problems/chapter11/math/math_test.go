package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	mn      float64
	mx      float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		m1 := Min(pair.values)
		m2 := Max(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v)
		}
		if m1 != pair.mn {
			t.Error("For", pair.values,
				"expected", pair.mn,
				"got", m1)
		}
		if m2 != pair.mx {
			t.Error("For", pair.values,
				"expected", pair.mx,
				"got", m2)
		}
	}
}
