package main

import "sort"

func findRadius(houses []int, heaters []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	sort.Ints(heaters)
	n := len(heaters)
	res := 0
	for _, house := range houses {
		left, right := 0, n
		for left < right {
			mid := left + (right-left)/2
			if heaters[mid] < house {
				left = mid + 1
			} else {
				right = mid
			}
		}
		dist1 := INT_MAX
		if right != n {
			dist1 = heaters[right] - house
		}
		dist2 := INT_MAX
		if right != 0 {
			dist2 = house - heaters[right-1]
		}
		d := dist1
		if dist2 < d {
			d = dist2
		}
		if d > res {
			res = d
		}
	}
	return res
}
