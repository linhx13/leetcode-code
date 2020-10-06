package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]
	if nums[0] > dp[1] {
		dp[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = nums[i] + dp[i-2]
		if dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 7, 9, 3, 1}
	// nums := []int{2, 1}
	nums := []int{}
	fmt.Println(rob(nums))
}
