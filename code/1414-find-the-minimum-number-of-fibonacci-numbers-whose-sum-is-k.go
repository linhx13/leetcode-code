package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	nums := []int{1, 1}
	for nums[len(nums)-1] < k {
		n := len(nums)
		nums = append(nums, nums[n-1]+nums[n-2])
	}
	res := 0
	for k > 0 {
		idx := search(nums, k)
		k -= nums[idx]
		res++
	}
	return res
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low - 1
}

func main() {
	k := 19
	fmt.Println(findMinFibonacciNumbers(k))
}
