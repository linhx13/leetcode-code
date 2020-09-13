package main

import "fmt"

func maxProduct(nums []int) int {
	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX
	res := INT_MIN
	pmax, pmin := 1, 1
	for i := 0; i < len(nums); i++ {
		cmax, cmin := 1, 1
		if nums[i] >= 0 {
			cmax = max(nums[i], pmax*nums[i])
			cmin = min(nums[i], pmin*nums[i])
		} else {
			cmax = max(nums[i], pmin*nums[i])
			cmin = min(nums[i], pmax*nums[i])
		}
		pmax, pmin = cmax, cmin
		if pmax > res {
			res = pmax
		}
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

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	// nums := []int{2, 3, -2, 4}
	// nums := []int{-2, 0, -1}
	// nums := []int{-2, -1, 0, 1}
	// nums := []int{-2, 3, -4}
	nums := []int{-4, -3, -2}
	fmt.Println(maxProduct(nums))
}
