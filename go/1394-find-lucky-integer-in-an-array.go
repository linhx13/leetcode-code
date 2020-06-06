package main

func findLucky(arr []int) int {
	cnts := make(map[int]int)
	for _, n := range arr {
		cnts[n] = cnts[n] + 1
	}
	res := 0
	for _, n := range arr {
		if n == cnts[n] {
			if n > res {
				res = n
			}
		}
	}
	return res
}
