package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	n := len(hand)
	if n%W != 0 {
		return false
	}
	counts := make(map[int]int)
	nums := []int{}
	for _, x := range hand {
		counts[x]++
		if counts[x] == 1 {
			nums = append(nums, x)
		}
	}
	sort.Ints(nums)
	i := 0
	for i < len(nums) {
		x := nums[i]
		if counts[x] == 0 {
			for i < len(nums) && counts[nums[i]] == 0 {
				i++
			}
			continue
		}
		for y := x + 1; y < x+W; y++ {
			if counts[y] == 0 {
				return false
			}
			counts[y]--
		}
		counts[x]--
	}
	return true
}

func main() {
	// hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	hand := []int{1, 2, 3, 4, 5, 6}
	W := 2
	fmt.Println(isNStraightHand(hand, W))
}
