package main

import "fmt"

func minCost(s string, cost []int) int {
	var last byte
	sum, max, cnt := 0, 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] != last && last != 0 {
			if cnt > 1 {
				res += sum - max
			}
			sum, max, cnt = 0, 0, 0
		}
		cnt++
		sum += cost[i]
		if cost[i] > max {
			max = cost[i]
		}
		last = s[i]
	}
	if cnt > 1 {
		res += sum - max
	}
	return res
}

func main() {
	// s := "abaac"
	// cost := []int{1, 2, 3, 4, 5}
	// s := "abc"
	// cost := []int{1, 2, 3}
	s := "aabaa"
	cost := []int{1, 2, 3, 4, 1}
	fmt.Println(minCost(s, cost))
}
