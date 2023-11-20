package split

import (
	"fmt"
)

type dynamicSolution struct {
	split     map[int]int
	bucketCnt int
	overflow  int
}

func newDynamicSolution() dynamicSolution {
	return dynamicSolution{
		split: make(map[int]int),
	}
}

func (d dynamicSolution) copy() dynamicSolution {
	newSolution := d
	newSolution.split = make(map[int]int)
	for k, v := range d.split {
		newSolution.split[k] = v
	}
	return newSolution
}

func filterMinOverflow(solutions []dynamicSolution) []dynamicSolution {
	if len(solutions) == 0 {
		return solutions
	}
	var filtered []dynamicSolution
	minOverflow := solutions[0].overflow
	for _, solution := range solutions {
		if solution.overflow == minOverflow {
			filtered = append(filtered, solution)
		}
		if solution.overflow < minOverflow {
			filtered = make([]dynamicSolution, 0)
			filtered = append(filtered, solution)
			minOverflow = solution.overflow
		}
	}
	return filtered
}

func filterOneMinBuckets(solutions []dynamicSolution) dynamicSolution {
	var filtered dynamicSolution
	minBuckets := solutions[0].bucketCnt
	for _, solution := range solutions {
		if solution.bucketCnt <= minBuckets {
			filtered = solution
			minBuckets = solution.overflow
		}
	}
	return filtered
}

func Dynamic(order int, bucketsArg []int, depth int) map[int]int {
	if order == 0 {
		return make(map[int]int)
	}

	buckets := make([]int, len(bucketsArg))
	copy(buckets, bucketsArg)

	solutions := dynamic(order, bucketsArg, dynamicSolution{}, depth)

	fmt.Printf("%v\n", len(solutions))

	minOver := filterMinOverflow(solutions)

	return filterOneMinBuckets(minOver).split
}

func dynamic(order int, bucketsArg []int, solution dynamicSolution, depth int) []dynamicSolution {
	if depth == 0 {
		return []dynamicSolution{}
	}

	solutions := make([]dynamicSolution, 0)
	for _, bucketSize := range bucketsArg {
		newSolution := solution.copy()
		decreasedOrder := order - bucketSize

		if _, ok := newSolution.split[bucketSize]; ok {
			newSolution.split[bucketSize] += 1
		} else {
			newSolution.split[bucketSize] = 1
		}
		newSolution.bucketCnt += 1
		if decreasedOrder <= 0 {
			newSolution.overflow = -decreasedOrder
			solutions = append(solutions, newSolution)
		} else {
			x := dynamic(decreasedOrder, bucketsArg, newSolution, depth-1)
			if len(x) == 0 {
				continue
			}
			solutions = append(solutions, x...)
		}
	}
	return solutions
}
