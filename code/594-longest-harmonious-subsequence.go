package main

import (
	"fmt"
)

func findLHS(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	count := make(map[int]int)
	res := 0
	for _, n := range nums {
		count[n]++
		v := count[n]
		if v1, ok := count[n+1]; ok && (v1+v > res) {
			res = v1 + v
		}
		if v2, ok := count[n-1]; ok && (v2+v > res) {
			res = v2 + v
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))
}
