package main

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sum = sum / k
	sums := make([]int, k)
	return helper(nums, sum, 0, sums)
}

func helper(nums []int, sum int, idx int, sums []int) bool {
	if idx == len(nums) {
		res := true
		for i := 0; i < len(sums); i++ {
			if sums[i] != sum {
				res = false
				break
			}
		}
		return res
	}
	res := false
	for i := 0; i < len(sums); i++ {
		if nums[idx]+sums[i] <= sum {
			sums[i] += nums[idx]
			res = res || helper(nums, sum, idx+1, sums)
			sums[i] -= nums[idx]
			if res {
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 2}
	k := 4
	fmt.Println(canPartitionKSubsets(nums, k))
}
