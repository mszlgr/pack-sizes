package split

import (
	"reflect"
	"testing"
)

func TestDynamic(t *testing.T) {
	cases := []struct {
		name      string
		orders    int
		packSizes []int
		expected  map[int]int
		depth     int
	}{
		{
			name:      "no orders",
			orders:    0,
			packSizes: []int{1, 2, 3},
			expected:  map[int]int{},
			depth:     8,
		},
		{
			name:      "one segment, fit",
			orders:    100,
			packSizes: []int{100},
			expected:  map[int]int{100: 1},
			depth:     8,
		},
		{
			name:      "two segments, by one order",
			orders:    101,
			packSizes: []int{100},
			expected:  map[int]int{100: 2},
			depth:     8,
		},
		{
			name:      "one of each",
			orders:    111111,
			packSizes: []int{1, 10, 100, 1000, 10000, 100000},
			expected:  map[int]int{1: 1, 10: 1, 100: 1, 1000: 1, 10000: 1, 100000: 1},
			depth:     8,
		},
		{
			name:      "diff of each",
			orders:    321,
			packSizes: []int{1, 10, 100, 1000, 10000, 100000},
			expected:  map[int]int{1: 1, 10: 2, 100: 3},
			depth:     8,
		},
		{
			name:      "dynamic is correct",
			orders:    45,
			packSizes: []int{15, 25},
			expected:  map[int]int{15: 3},
			depth:     8,
		},
	}

	for _, c := range cases {
		ret := Dynamic(c.orders, c.packSizes, Options{Depth: c.depth})
		if !reflect.DeepEqual(ret, c.expected) {
			t.Errorf("in test: %s, expected %+v, got %+v", c.name, c.expected, ret)
		}
	}

}
