package math

import "testing"

type testpair struct {
	values  []int
	average int
}

var tests = []testpair{
	{[]int{100, 50}, 75},
	{[]int{60, 60}, 60},
	{[]int{6, 50}, 28},

}

func TestAverage(t *testing.T) {
	//this is test for func Average, all test receive *testing.T as input
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For",pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
