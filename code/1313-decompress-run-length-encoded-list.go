package main

import "fmt"

func decompressRLElist(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}

func main() {
	// nums := []int{1, 1, 2, 3}
	// nums := []int{1, 2, 3, 4}
	nums := []int{1, 1}
	fmt.Println(decompressRLElist(nums))
}
