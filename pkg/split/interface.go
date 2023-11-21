package split

type Split func(order int, packSizes []int, opt Options) map[int]int
