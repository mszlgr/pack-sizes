package split

import (
	"reflect"
	"testing"
)

func TestGreedy(t *testing.T) {
	cases := []struct {
		name      string
		orders    int
		packSizes []int
		expected  map[int]int
	}{
		{
			name:      "no orders",
			orders:    0,
			packSizes: []int{1, 2, 3},
			expected:  map[int]int{},
		},
		{
			name:      "one segment, fit",
			orders:    100,
			packSizes: []int{100},
			expected:  map[int]int{100: 1},
		},
		{
			name:      "two segments, by one order",
			orders:    101,
			packSizes: []int{100},
			expected:  map[int]int{100: 2},
		},
		{
			name:      "one of each",
			orders:    111111,
			packSizes: []int{1, 10, 100, 1000, 10000, 100000},
			expected:  map[int]int{1: 1, 10: 1, 100: 1, 1000: 1, 10000: 1, 100000: 1},
		},
		{
			name:      "diff of each",
			orders:    123456,
			packSizes: []int{1, 10, 100, 1000, 10000, 100000},
			expected:  map[int]int{1: 6, 10: 5, 100: 4, 1000: 3, 10000: 2, 100000: 1},
		},
		{
			name:      "binary split",
			orders:    135,
			packSizes: []int{1, 2, 4, 8, 16, 32, 64, 128, 256},
			expected:  map[int]int{1: 1, 2: 1, 4: 1, 128: 1},
		},
		{
			name:      "greedy is wrong",
			orders:    45,
			packSizes: []int{15, 25},
			expected:  map[int]int{15: 2, 25: 1},
		},
	}

	for _, c := range cases {
		ret := Greedy(c.orders, c.packSizes, Options{})
		if !reflect.DeepEqual(ret, c.expected) {
			t.Errorf("in test: %s, expected %+v, got %+v", c.name, c.expected, ret)
		}
	}

}
