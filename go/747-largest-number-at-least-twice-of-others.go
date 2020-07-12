package main

func dominantIndex(nums []int) int {
	max, maxIdx := -1, -1
	for idx, n := range nums {
		if n > max {
			max, maxIdx = n, idx
		}
	}
	for _, n := range nums {
		if n != max && max < 2*n {
			return -1
		}
	}
	return maxIdx
}
