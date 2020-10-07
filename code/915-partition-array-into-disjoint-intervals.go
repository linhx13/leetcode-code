package main

import "fmt"

func partitionDisjoint(A []int) int {
	max := make([]int, len(A))
	min := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		min[i] = int(^uint(0) >> 1)
		max[i] = 0
	}
	for i := 0; i < len(A); i++ {
		if i == 0 {
			max[i] = A[i]
		} else {
			if A[i] > max[i-1] {
				max[i] = A[i]
			} else {
				max[i] = max[i-1]
			}
		}
	}
	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			min[i] = A[i]
		} else {
			if A[i] < min[i+1] {
				min[i] = A[i]
			} else {
				min[i] = min[i+1]
			}
		}
	}
	for i := 0; i < len(A); i++ {
		if i+1 < len(A) && max[i] <= min[i+1] {
			return i + 1
		}
	}
	return 0
}

func main() {
	A := []int{5, 0, 3, 8, 6}
	// A := []int{1, 1, 1, 0}
	fmt.Println(partitionDisjoint(A))
}
