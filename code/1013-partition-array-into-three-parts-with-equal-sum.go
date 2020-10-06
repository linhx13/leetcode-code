package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	s := 0
	for _, i := range A {
		s += i
	}

	if s%3 != 0 {
		return false
	}
	var s1 int = s / 3
	var s2 int = 2 * s1

	ind1, ind2 := -1, -1
	preSum := 0
	for i := 0; i < len(A)-1; i++ {
		preSum += A[i]

		if preSum == s1 && ind1 == -1 {
			ind1 = i
		} else if preSum == s2 && ind1 != -1 {
			ind2 = i
			break
		}
	}
	return ind1 != -1 && ind2 != -1
}

func main() {
	// A := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	// A := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	// A := []int{6, 1, 1, 13, -1, 0, -10, 20}
	A := []int{29, 31, 27, -10, -67, 22, 15, -1, -16, -29, 59, -18, 48}
	// A := []int{1, -1, 1, -1}
	fmt.Println(canThreePartsEqualSum(A))
}
