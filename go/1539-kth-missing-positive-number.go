package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	x := arr[len(arr)-1] - len(arr)
	if k > x {
		return arr[len(arr)-1] + k - x
	}
	m := make(map[int]bool)
	for _, n := range arr {
		m[n] = true
	}
	for n := 1; n > 0; n++ {
		if !m[n] {
			k--
			if k == 0 {
				return n
			}
		}
	}
	return 0
}

func main() {
	// arr := []int{2, 3, 4, 7, 11}
	// k := 5
	arr := []int{1, 2, 3, 4}
	k := 1000
	fmt.Println(findKthPositive(arr, k))
}
