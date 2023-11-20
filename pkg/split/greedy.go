package split

import (
	"sort"
)

func Greedy(order int, bucketsArg []int) map[int]int {
	segments := make(map[int]int)
	buckets := make([]int, len(bucketsArg))
	copy(buckets, bucketsArg)

	sort.Sort(sort.Reverse(sort.IntSlice(buckets)))
	for _, size := range buckets {
		segmentCnt := order / size
		order = order - size*segmentCnt
		if segmentCnt == 0 {
			continue
		}
		segments[size] = segmentCnt
	}
	if order > 0 {
		lastPackSize := buckets[len(buckets)-1]
		segments[lastPackSize] = segments[lastPackSize] + 1
	}
	return segments
}
