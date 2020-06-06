package main

import "math"

func shortestToChar(S string, C byte) []int {
	const INT_MAX = int(^uint(0) >> 1)
	res := []int{}
	for i := 0; i < len(S); i++ {
		res = append(res, INT_MAX)
	}
	idx := -1
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			idx = i
		}
		if idx < 0 {
			continue
		}
		res[i] = int(math.Abs(float64(i - idx)))
	}
	idx = -1
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			idx = i
		}
		if idx < 0 {
			continue
		}
		abs := int(math.Abs(float64(i - idx)))
		if abs < res[i] {
			res[i] = abs
		}
	}
	return res
}
