package split

type Split func(order int, packSizes []int) map[int]int
