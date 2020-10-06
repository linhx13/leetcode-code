package main

import (
	"fmt"
	"math"
)

const INT_MAX = int(^uint(0) >> 1)

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	ds := make([]int, n)
	for i := 0; i < n; i++ {
		ds[i] = INT_MAX
	}
	for i := 1; i < n; i++ {
		ds[i] = dist(points[0], points[i])
	}
	res := 0
	for i := 1; i < n; i++ {
		minIdx := minElement(ds)
		res += ds[minIdx]
		ds[minIdx] = INT_MAX
		for j := 0; j < n; j++ {
			if ds[j] == INT_MAX {
				continue
			}
			ds[j] = min(ds[j], dist(points[j], points[minIdx]))
		}
	}
	return res
}

func dist(x, y []int) int {
	return int(math.Abs(float64(x[0]-y[0]))) + int(math.Abs(float64(x[1]-y[1])))
}

func minElement(nums []int) int {
	min := INT_MAX
	idx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			idx = i
		}
	}
	return idx
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	points := [][]int{[]int{0, 0}, []int{2, 2}, []int{3, 10}, []int{5, 2}, []int{7, 0}}
	fmt.Println(minCostConnectPoints(points))
}
