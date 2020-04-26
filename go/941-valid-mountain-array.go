package main

import "fmt"

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	idx := 1
	for ; idx < len(A)-1; idx++ {
		if !(A[idx] > A[idx-1]) {
			if idx == 1 {
				return false
			}
			break
		}
	}
	for ; idx < len(A)-1; idx++ {
		if !(A[idx+1] < A[idx]) {
			return false
		}
	}
	return A[idx] < A[idx-1]
}

func main() {
	// b := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// b := []int{1, 3, 2}
	b := []int{3, 5, 5}
	fmt.Println(validMountainArray(b))
}
