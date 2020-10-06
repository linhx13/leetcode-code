package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	high := -1
	for i := 0; i < len(nums); i++ {
		n, idx := nums[i], index[i]
		if idx > high {
			res[idx] = n
			high = idx
		} else {
			for j := high; j >= idx; j-- {
				res[j+1] = res[j]
			}
			res[idx] = n
			high++
		}
	}
	return res
}

func main() {
	// nums := []int{0, 1, 2, 3, 4}
	// index := []int{0, 1, 2, 2, 1}
	// nums := []int{1, 2, 3, 4, 0}
	// index := []int{0, 1, 2, 3, 0}
	nums := []int{1}
	index := []int{0}
	fmt.Println(createTargetArray(nums, index))
}
