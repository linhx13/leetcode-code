package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	res := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] > nums[k] {
					res++
				}
				if nums[i]+nums[j] <= nums[k] {
					break
				}
			}
		}
	}
	return res
}

func main() {
	// nums := []int{2, 2, 3, 4}
	nums := []int{3, 1, 1}
	fmt.Println(triangleNumber(nums))
}
