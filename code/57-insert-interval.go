package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	start := 0
	inStart := false
	for ; start < len(intervals); start++ {
		if intervals[start][0] <= newInterval[0] && newInterval[0] <= intervals[start][1] {
			inStart = true
			break
		}
		if newInterval[0] < intervals[start][0] {
			inStart = false
			break
		}
	}
	end := start
	inEnd := false
	for ; end < len(intervals); end++ {
		if intervals[end][0] <= newInterval[1] && newInterval[1] <= intervals[end][1] {
			inEnd = true
			break
		}
		if newInterval[1] < intervals[end][0] {
			inEnd = false
			break
		}
	}
	res := [][]int{}
	res = append(res, intervals[0:start]...)
	x := []int{newInterval[0], newInterval[1]}
	if start >= len(intervals) {
		res = append(res, x)
	} else {
		if inStart {
			x[0] = intervals[start][0]
		}
		if inEnd {
			x[1] = intervals[end][1]
			res = append(res, x)
			res = append(res, intervals[end+1:len(intervals)]...)
		} else {
			res = append(res, x)
			res = append(res, intervals[end:len(intervals)]...)
		}
	}
	return res
}

func main() {
	// intervals := [][]int{[]int{1, 3}, []int{6, 9}}
	// newInterval := []int{2, 5}
	intervals := [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
}
