package main

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	judge := true
	for i := 0; i < K; i++ {
		if A[i] < 0 {
			A[i] = -A[i]
		} else {
			if (K-i)%2 == 0 {
				judge = true
			} else {
				judge = false
			}
			break
		}
	}
	sum := 0
	min := int(^uint(0) >> 1)
	for _, a := range A {
		sum += a
		if a < min {
			min = a
		}
	}
	if judge {
		return sum
	} else {
		return sum - 2*min
	}
}
