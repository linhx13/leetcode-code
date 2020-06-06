package main

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, c := range chips {
		if c%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd > even {
		return even
	} else {
		return odd
	}
}
