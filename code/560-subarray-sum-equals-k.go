package main

import "fmt"

func subarraySum(nums []int, k int) int {
	m := make(map[int][]int)
	sums := make([]int, len(nums))
	m[0] = append(m[0], -1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[i] = nums[i]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
		m[sums[i]] = append(m[sums[i]], i)
	}
	fmt.Println(m)
	res := 0
	for i := 0; i < len(nums); i++ {
		x := sums[i] - k
		for _, idx := range m[x] {
			if idx < i {
				res++
			} else {
				break
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
