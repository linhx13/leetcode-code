package main

import "fmt"

func numSubarraysWithSum(A []int, S int) int {
	m := make(map[int][]int)
	sums := make([]int, len(A))
	m[0] = append(m[0], -1)
	for i := 0; i < len(A); i++ {
		if i == 0 {
			sums[i] = A[i]
		} else {
			sums[i] = sums[i-1] + A[i]
		}
		m[sums[i]] = append(m[sums[i]], i)
	}
	res := 0
	for i := 0; i < len(A); i++ {
		x := sums[i] - S
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
	A := []int{1, 0, 1, 0, 1}
	S := 2
	// A := []int{0, 0, 0, 0, 0}
	// S := 0
	fmt.Println(numSubarraysWithSum(A, S))
}
