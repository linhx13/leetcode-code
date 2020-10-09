package main

import "fmt"

func lengthOfLIS(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
