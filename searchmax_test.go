package search

import (
	"testing"
)

func TestMax(t *testing.T) {

	type testStructSlice struct {
		a string
		b int
	}

	IntSlice := []interface{}{1, 2, 3, 4, 5, 6, 11, 10, 9, 18, 7}
	FloatSlice := []interface{}{1.0, 2.2, 0.3, 0.4, 5.0, 6.0, 11.1, 10.4, 9.0, 1.8, 0.7}

	StructSlice := []interface{}{
		testStructSlice{"AAA", 1},
		testStructSlice{"DDD", 2},
		testStructSlice{"KKK", 3},
		testStructSlice{"GGG", 4},
		testStructSlice{"BBB", 5},
	}

	cases := []struct {
		arr        []interface{}
		maxIndex   int
		comparator MaxComparator
	}{
		{IntSlice, 9, func(i, max int) bool { return IntSlice[i].(int) > IntSlice[max].(int) }},
		{FloatSlice, 6, func(i, max int) bool { return FloatSlice[i].(float64) > FloatSlice[max].(float64) }},
		{StructSlice, 2, func(i, max int) bool {
			return StructSlice[i].(testStructSlice).a > StructSlice[max].(testStructSlice).a
		}},
	}

	for _, testCase := range cases {

		m := Max(testCase.arr, testCase.comparator)
		if m != testCase.arr[testCase.maxIndex] {
			t.Errorf("Got %v Expected %v", m, testCase.arr[testCase.maxIndex])
		}
	}
}
